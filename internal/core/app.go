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
	return &AppImpl{
		Port:   port,
		Router: router,
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

		// Handle each connection in a goroutine
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
		result := router.Handle(command, args[1:])
		conn.Write(result)
		if command != "INFO" {
			fmt.Println("result", string(result))
		}
	}
}
