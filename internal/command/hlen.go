package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HLen(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'hlen' command")
	}

	key := args[0]
	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		result, err := utils.EncodeRESP(0)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)
	length := len(hash)

	result, err := utils.EncodeRESP(length)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
