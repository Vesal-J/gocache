package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HMGet(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'hmget' command")
	}

	key := args[0]
	fields := args[1:]

	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESPArray([]string{})
	}

	if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)

	response := make([]string, 0, len(fields))
	for _, field := range fields {
		value, exists := hash[field]
		if exists {
			response = append(response, value)
		} else {
			response = append(response, "")
		}
	}

	return utils.ToRESPArray(response)
}
