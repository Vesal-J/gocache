package command

import (
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HSet(args []string) []byte {
	if len(args) < 3 || len(args)%2 != 1 {
		return utils.ToRESPError("wrong number of arguments for 'hset' command")
	}

	key := args[0]

	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		cacheObj = store.CacheObject{
			Type:      store.HASH,
			Value:     make(map[string]string),
			TTL:       0,
			CreatedAt: time.Now().Unix(),
		}
	} else if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)

	for i := 1; i < len(args); i += 2 {
		field := args[i]
		value := args[i+1]
		hash[field] = value
	}

	c.Store.Caches[key] = cacheObj

	return utils.ToRESP("OK")
}
