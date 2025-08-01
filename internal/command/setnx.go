package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) SetNX(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'setnx' command")
	}

	key := args[0]
	value := args[1]

	if _, exists := c.Store.Caches[key]; exists {
		return utils.ToRESP("(integer) 0")
	}

	c.Set([]string{key, value})
	return utils.ToRESP("(integer) 1")
}
