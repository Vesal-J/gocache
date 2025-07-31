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
}

type CommandImpl struct {
	Store *store.Store
}

func NewCommand(store *store.Store) Command {
	return &CommandImpl{
		Store: store,
	}
}
