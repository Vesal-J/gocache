package core

import (
	"strings"

	"github.com/vesal-j/gocache/internal/command"
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

type Router struct {
	Store      store.Store
	Command    command.Command
	CommandMap map[string]func([]string) []byte
}

func NewRouter(store *store.Store, command command.Command) *Router {
	router := &Router{
		Store:   *store,
		Command: command,
	}

	router.CommandMap = map[string]func([]string) []byte{
		"ping":   command.Ping,
		"get":    command.Get,
		"set":    command.Set,
		"client": command.Client,
		"info":   command.Info,
		"auth":   command.Auth,
		"exists": command.Exists,
		"ttl":    command.TTL,
		"type":   command.Type,
		"dbsize": command.DBSIZE,
		"memory": command.Memory,
		"scan":   command.Scan,
		"config": command.Config,
	}

	return router
}

func (r *Router) Handle(command string, args []string) []byte {
	command = strings.TrimSpace(strings.ToLower(command))

	if command == "" {
		return utils.ToRESP("ERR empty command")
	}

	commandFunc, exists := r.CommandMap[command]
	if !exists {
		return utils.ToRESP("ERR unknown command")
	}

	return commandFunc(args)
}
