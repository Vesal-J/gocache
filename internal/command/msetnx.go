package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) MSetNX(args []string) []byte {
	if len(args)%2 != 0 {
		return utils.ToRESPError("wrong number of arguments for 'msetnx' command")
	}

	for i := 0; i < len(args); i += 2 {
		key := args[i]
		value := args[i+1]

		if _, exists := c.Store.Caches[key]; exists {
			continue
		}

		c.Set([]string{key, value})
	}

	return utils.ToRESP("OK")
}
