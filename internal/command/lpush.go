package command

import (
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) LPush(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'lpush' command")
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

		newList := make([]string, 0, len(values)+len(list))
		newList = append(newList, values...)
		newList = append(newList, list...)

		c.Store.Caches[key] = store.CacheObject{
			Type:      store.LIST,
			Value:     newList,
			TTL:       existingObj.TTL,
			CreatedAt: existingObj.CreatedAt,
		}

		result, err := utils.EncodeRESP(len(newList))
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
			TTL:       0,
			CreatedAt: time.Now().Unix(),
		}

		result, err := utils.EncodeRESP(len(newList))
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}
}
