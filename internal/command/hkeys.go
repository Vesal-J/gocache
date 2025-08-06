package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HKeys(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'hkeys' command")
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

	keys := make([]string, 0, len(hash))
	for field := range hash {
		keys = append(keys, field)
	}

	return utils.ToRESPArray(keys)
}
