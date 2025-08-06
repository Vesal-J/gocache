package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HExists(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'hexists' command")
	}

	key := args[0]
	field := args[1]

	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("0")
	}

	if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)

	if _, fieldExists := hash[field]; fieldExists {
		return utils.ToRESP("1")
	}

	return utils.ToRESP("0")
}
