package command

import "github.com/vesal-j/gocache/internal/store"

type Command interface {
	Set(args []string) string
	Get(args []string) string
}

type CommandImpl struct {
	Store *store.Store
}

func NewCommand(store *store.Store) *CommandImpl {
	return &CommandImpl{
		Store: store,
	}
}
