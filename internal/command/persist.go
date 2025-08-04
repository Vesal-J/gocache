package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Persist(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'persist' command")
	}

	key := args[0]
	cache, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("(integer) 0")
	}

	cache.TTL = 0
	c.Store.Caches[key] = cache

	return utils.ToRESP("(integer) 1")
}
