package core

import (
	"strings"

	"github.com/vesal-j/gocache/internal/command"
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

type Router struct {
	Store   store.Store
	Command command.Command
}

func NewRouter(store *store.Store, command command.Command) *Router {
	return &Router{
		Store:   *store,
		Command: command,
	}
}

func (r *Router) Handle(command string, args []string) []byte {
	command = strings.TrimSpace(strings.ToLower(command))
	switch command {
	case "ping":
		return r.Command.Ping(args)
	case "get":
		return r.Command.Get(args)
	case "set":
		return r.Command.Set(args)
	case "client":
		return r.Command.Client(args)
	case "info":
		return r.Command.Info(args)
	case "auth":
		return r.Command.Auth(args)
	case "exists":
		return r.Command.Exists(args)
	case "ttl":
		return r.Command.TTL(args)
	case "type":
		return r.Command.Type(args)
	case "dbsize":
		return r.Command.DBSIZE(args)
	case "memory":
		return r.Command.Memory(args)
	case "scan":
		return r.Command.Scan(args)
	case "config":
		return r.Command.Config(args)
	default:
		return utils.ToRESP("ERR unknown command")
	}
}
