package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HIncrByFloat(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'hincrbyfloat' command")
	}

	key := args[0]
	field := args[1]
	incrementStr := args[2]

	increment, err := strconv.ParseFloat(incrementStr, 64)
	if err != nil {
		return utils.ToRESPError("value is not a valid float")
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

	currentFloat, err := strconv.ParseFloat(currentValue, 64)
	if err != nil {
		return utils.ToRESPError("hash value is not a valid float")
	}

	newValue := currentFloat + increment
	hash[field] = strconv.FormatFloat(newValue, 'f', -1, 64)

	c.Store.Caches[key] = cacheObj

	return utils.ToRESP(strconv.FormatFloat(newValue, 'f', -1, 64))
}
