package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) LSet(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'lset' command")
	}

	key := args[0]
	index, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid index")
	}

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

	if index < 0 {
		index = len(list) + index
	}

	if index < 0 || index >= len(list) {
		return utils.ToRESPError("index out of range")
	}

	list[index] = args[2]

	c.Store.Caches[key] = store.CacheObject{
		Type:      store.LIST,
		Value:     list,
		TTL:       existingObj.TTL,
		CreatedAt: existingObj.CreatedAt,
	}

	return utils.ToRESP("OK")
}
