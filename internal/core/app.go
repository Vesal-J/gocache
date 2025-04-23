package core

import (
	"fmt"
	"net"
	"strings"
)

type App interface {
	Listen()
	handleConnection(net.Conn)
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

func (l *AppImpl) Listen() {
	fmt.Printf("Starting server on port %d\n", l.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", l.Port))
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}

		fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())

		// Handle connection in a goroutine
		go l.handleConnection(conn)
	}
}

func (a *AppImpl) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Connection closed: %v\n", err)
			return
		}

		input := strings.TrimSpace(string(buffer[:n]))
		input = strings.Trim(input, "{}")

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		args := parts[1:]

		result := a.Router.Handle(command, args)

		response := fmt.Sprintf("%s\n", result)
		conn.Write([]byte(response))
	}
}
