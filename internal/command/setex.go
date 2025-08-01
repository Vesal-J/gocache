package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) SetEX(args []string) []byte {
	if len(args) < 3 {
		return utils.ToRESPError("wrong number of arguments for 'setex' command")
	}

	key := args[0]
	value := args[2]
	ttl, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid ttl")
	}

	c.Store.Caches[key] = store.CacheObject{
		Key:       key,
		Value:     value,
		TTL:       time.Duration(ttl) * time.Second,
		CreatedAt: time.Now().Unix(),
	}

	return utils.ToRESP("OK")
}
