package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Expire(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'expire' command")
	}

	key := args[0]
	ttl, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid ttl")
	}

	cache, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("(integer) 0")
	}

	cache.TTL = time.Duration(ttl) * time.Second
	cache.CreatedAt = time.Now().Unix()
	c.Store.Caches[key] = cache

	return utils.ToRESP("(integer) 1")
}
