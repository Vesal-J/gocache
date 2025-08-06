package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HIncrBy(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'hincrby' command")
	}

	key := args[0]
	field := args[1]
	incrementStr := args[2]

	increment, err := strconv.ParseInt(incrementStr, 10, 64)
	if err != nil {
		return utils.ToRESPError("value is not an integer or out of range")
	}

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

	currentValue := "0"
	if existingValue, exists := hash[field]; exists {
		currentValue = existingValue
	}

	currentInt, err := strconv.ParseInt(currentValue, 10, 64)
	if err != nil {
		return utils.ToRESPError("hash value is not an integer")
	}

	newValue := currentInt + increment
	hash[field] = strconv.FormatInt(newValue, 10)

	c.Store.Caches[key] = cacheObj

	return utils.ToRESP(strconv.FormatInt(newValue, 10))
}
