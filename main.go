package main

import (
	"github.com/vesal-j/gocache/internal/command"
	"github.com/vesal-j/gocache/internal/core"
	"github.com/vesal-j/gocache/internal/store"
)

func main() {
	store := store.NewStore()
	command := command.NewCommand(store)
	router := core.NewRouter(store, &command)
	app := core.NewApp(6380, router)
	app.Listen()
}
