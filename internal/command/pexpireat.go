package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) PExpireAt(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'pexpireat' command")
	}

	key := args[0]
	ttl := args[1]

	cache, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("(integer) 0")
	}

	ttlInt, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return utils.ToRESPError("invalid ttl")
	}

	cache.TTL = time.Duration(ttlInt) * time.Millisecond
	c.Store.Caches[key] = cache

	return utils.ToRESP("(integer) 1")
}
