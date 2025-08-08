package command

import (
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) RPush(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'rpush' command")
	}

	key := args[0]
	values := args[1:]

	existingObj, exists := c.Store.Caches[key]

	if exists {
		if existingObj.Type != store.LIST {
			return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
		}

		list, ok := existingObj.Value.([]string)
		if !ok {
			list = []string{}
		}

		list = append(list, values...)

		c.Store.Caches[key] = store.CacheObject{
			Type:      store.LIST,
			Value:     list,
			TTL:       existingObj.TTL,
			CreatedAt: existingObj.CreatedAt,
		}

		result, err := utils.EncodeRESP(len(list))
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	} else {
		newList := make([]string, len(values))
		copy(newList, values)

		c.Store.Caches[key] = store.CacheObject{
			Type:      store.LIST,
			Value:     newList,
			TTL:       0, // No TTL for new list
			CreatedAt: time.Now().Unix(),
		}

		result, err := utils.EncodeRESP(len(newList))
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}
}
