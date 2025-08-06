package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HGetAll(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'hgetall' command")
	}

	key := args[0]
	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESPArray([]string{})
	}

	if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)

	response := make([]string, 0, len(hash)*2)
	for field, value := range hash {
		response = append(response, field, value)
	}

	return utils.ToRESPArray(response)
}
