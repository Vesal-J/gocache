package command

import (
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HMSet(args []string) []byte {
	if len(args) < 3 || len(args)%2 != 1 {
		return utils.ToRESPError("wrong number of arguments for 'hmset' command")
	}

	key := args[0]
	hash := make(map[string]string)

	for i := 1; i < len(args); i += 2 {
		hash[args[i]] = args[i+1]
	}

	cacheObj := store.CacheObject{
		Type:      store.HASH,
		Value:     hash,
		TTL:       0,
		CreatedAt: time.Now().Unix(),
	}

	c.Store.Caches[key] = cacheObj
	return utils.ToRESP("OK")
}
