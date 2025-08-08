package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) LRem(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'lrem' command")
	}

	key := args[0]
	count, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid count")
	}

	value := args[2]

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

	removedCount := 0

	if count > 0 {
		// Remove from left to right (head to tail)
		for i := 0; i < len(list) && removedCount < count; {
			if list[i] == value {
				// Remove element at index i
				list = append(list[:i], list[i+1:]...)
				removedCount++
			} else {
				i++
			}
		}
	} else if count < 0 {
		// Remove from right to left (tail to head)
		for i := len(list) - 1; i >= 0 && removedCount < -count; {
			if list[i] == value {
				// Remove element at index i
				list = append(list[:i], list[i+1:]...)
				removedCount++
			}
			i--
		}
	} else {
		// count == 0, remove all occurrences
		for i := 0; i < len(list); {
			if list[i] == value {
				// Remove element at index i
				list = append(list[:i], list[i+1:]...)
				removedCount++
			} else {
				i++
			}
		}
	}

	// If list becomes empty, delete the key (Redis behavior)
	if len(list) == 0 {
		delete(c.Store.Caches, key)
	} else {
		// Update the cache object
		c.Store.Caches[key] = store.CacheObject{
			Type:      store.LIST,
			Value:     list,
			TTL:       existingObj.TTL,
			CreatedAt: existingObj.CreatedAt,
		}
	}

	result, err := utils.EncodeRESP(removedCount)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
