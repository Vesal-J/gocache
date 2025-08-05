package core

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type App interface {
	Listen()
}

type AppImpl struct {
	Port   int
	Router *Router
}

func NewApp(port int, router *Router) App {
	app := &AppImpl{
		Port:   port,
		Router: router,
	}

	go app.startEventProcessor()
	return app
}

func (a *AppImpl) startEventProcessor() {
	for event := range a.Router.EventLoop {
		result := a.Router.Handle(event.Command, event.Args)
		(*event.Conn).Write(result)
		if event.Command != "INFO" {
			fmt.Println("result", string(result))
		}
	}
}

func (a *AppImpl) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis mock listening on port", a.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go handleConnection(conn, a.Router)
	}
}

func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "*") {
			continue
		}

		argCount, err := strconv.Atoi(strings.TrimPrefix(line, "*"))
		if err != nil || argCount == 0 {
			continue
		}

		args := make([]string, 0, argCount)

		for i := 0; i < argCount; i++ {
			dollarLine, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			dollarLine = strings.TrimSpace(dollarLine)
			if !strings.HasPrefix(dollarLine, "$") {
				continue
			}

			argLen, err := strconv.Atoi(strings.TrimPrefix(dollarLine, "$"))
			if err != nil || argLen < 0 {
				continue
			}

			argBytes := make([]byte, argLen)
			_, err = reader.Read(argBytes)
			if err != nil {
				return
			}

			_, err = reader.ReadString('\n')
			if err != nil {
				return
			}

			args = append(args, string(argBytes))
		}

		if len(args) == 0 {
			conn.Write([]byte("-ERR empty command\r\n"))
			continue
		}

		command := strings.ToUpper(args[0])
		fmt.Println(command, args)

		router.EventLoop <- Event{
			Command: command,
			Args:    args[1:],
			Conn:    &conn,
		}
	}
}
