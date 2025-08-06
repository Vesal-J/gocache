package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HDel(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'hdel' command")
	}

	key := args[0]
	fields := args[1:]

	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("(integer) 0")
	}

	if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)
	deleted := 0

	for _, field := range fields {
		if _, exists := hash[field]; exists {
			delete(hash, field)
			deleted++
		}
	}

	c.Store.Caches[key] = cacheObj
	return utils.ToRESP(strconv.Itoa(deleted))
}
