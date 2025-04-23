package core

import (
	"strings"

	"github.com/vesal-j/gocache/internal/command"
	"github.com/vesal-j/gocache/internal/store"
)

type Router struct {
	Store   store.Store
	Command *command.CommandImpl
}

func NewRouter(store *store.Store, command *command.CommandImpl) *Router {
	return &Router{
		Store:   *store,
		Command: command,
	}
}

func (r *Router) Handle(command string, args []string) string {
	command = strings.TrimSpace(strings.ToLower(command))

	switch command {
	case "get":
		return r.Command.Get(args)
	case "set":
		return r.Command.Set(args)
	default:
		return "ERR unknown command"
	}
}
