package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) ExpireAt(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'expireat' command")
	}

	key := args[0]
	timestamp, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return utils.ToRESPError("invalid timestamp")
	}

	cache, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("(integer) 0")
	}

	now := time.Now().Unix()
	ttlSeconds := timestamp - now

	// If timestamp is in the past or present, delete the key (expire immediately)
	if ttlSeconds <= 0 {
		delete(c.Store.Caches, key)
		return utils.ToRESP("(integer) 1")
	}

	// If timestamp is in the future, set positive TTL
	cache.TTL = time.Duration(ttlSeconds) * time.Second
	cache.CreatedAt = now
	c.Store.Caches[key] = cache

	return utils.ToRESP("(integer) 1")
}
