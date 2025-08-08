package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) PrPop(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'prpop' command")
	}

	key := args[0]
	existingObj, exists := c.Store.Caches[key]

	if !exists {
		return utils.ToRESPError("key does not exist")
	}

	if existingObj.Type != store.LIST {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	list, ok := existingObj.Value.([]string)
	if !ok {
		return utils.ToRESPError("value is not a list")
	}

	if len(list) == 0 {
		return utils.ToRESPError("list is empty")
	}

	popped := list[len(list)-1]
	list = list[:len(list)-1]

	// If list becomes empty, delete the key (Redis behavior)
	if len(list) == 0 {
		delete(c.Store.Caches, key)
	} else {
		c.Store.Caches[key] = store.CacheObject{
			Type:      store.LIST,
			Value:     list,
			TTL:       existingObj.TTL,
			CreatedAt: existingObj.CreatedAt,
		}
	}

	return utils.ToRESP(popped)
}
