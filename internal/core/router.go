package core

import (
	"net"
	"strings"

	"github.com/vesal-j/gocache/internal/command"
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

type Router struct {
	Store      store.Store
	Command    command.Command
	CommandMap map[string]func([]string) []byte
	EventLoop  chan Event
}

type Event struct {
	Command string
	Args    []string
	Conn    *net.Conn
}

func NewRouter(store *store.Store, command command.Command) *Router {
	router := &Router{
		Store:     *store,
		Command:   command,
		EventLoop: make(chan Event, 1000),
	}

	router.CommandMap = map[string]func([]string) []byte{
		"ping":      command.Ping,
		"get":       command.Get,
		"set":       command.Set,
		"client":    command.Client,
		"info":      command.Info,
		"auth":      command.Auth,
		"exists":    command.Exists,
		"ttl":       command.TTL,
		"type":      command.Type,
		"dbsize":    command.DBSIZE,
		"memory":    command.Memory,
		"scan":      command.Scan,
		"config":    command.Config,
		"strlen":    command.Strlen,
		"getrange":  command.Getrange,
		"command":   command.Command,
		"del":       command.Del,
		"incr":      command.Incr,
		"decr":      command.Decr,
		"incrby":    command.IncrBy,
		"decrby":    command.DecrBy,
		"append":    command.Append,
		"setex":     command.SetEX,
		"setnx":     command.SetNX,
		"psetex":    command.PSetEX,
		"mset":      command.MSet,
		"mget":      command.MGet,
		"msetnx":    command.MSetNX,
		"getset":    command.GetSet,
		"expire":    command.Expire,
		"expireat":  command.ExpireAt,
		"persist":   command.Persist,
		"pexpire":   command.PExpire,
		"pexpireat": command.PExpireAt,
		"keys":      command.Keys,
		"rename":    command.Rename,
		"renamenx":  command.RenameNX,
	}

	return router
}

func (r *Router) Handle(command string, args []string) []byte {
	command = strings.TrimSpace(strings.ToLower(command))

	if command == "" {
		return utils.ToRESPError("empty command")
	}

	commandFunc, exists := r.CommandMap[command]
	if !exists {
		return utils.ToRESPError("unknown command")
	}

	return commandFunc(args)
}
