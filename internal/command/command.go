package command

import "github.com/vesal-j/gocache/internal/store"

type Command interface {
	Auth(args []string) []byte
	Client(args []string) []byte
	Config(args []string) []byte
	DBSIZE(args []string) []byte
	Exists(args []string) []byte
	Get(args []string) []byte
	Info(args []string) []byte
	Memory(args []string) []byte
	Ping(args []string) []byte
	Scan(args []string) []byte
	Set(args []string) []byte
	TTL(args []string) []byte
	Type(args []string) []byte
	Strlen(args []string) []byte
	Getrange(args []string) []byte
	Command(args []string) []byte
	Del(args []string) []byte
	Incr(args []string) []byte
	IncrBy(args []string) []byte
	Decr(args []string) []byte
	DecrBy(args []string) []byte
	Append(args []string) []byte
	SetEX(args []string) []byte
	SetNX(args []string) []byte
	PSetEX(args []string) []byte
}

type CommandImpl struct {
	Store *store.Store
}

func NewCommand(store *store.Store) Command {
	return &CommandImpl{
		Store: store,
	}
}
