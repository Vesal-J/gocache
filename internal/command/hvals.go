package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HVals(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'hvals' command")
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

	values := make([]string, 0, len(hash))
	for _, value := range hash {
		values = append(values, value)
	}

	return utils.ToRESPArray(values)
}
