package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) BLPop(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'blpop' command")
	}

	// Parse timeout
	timeout, err := strconv.Atoi(args[len(args)-1])
	if err != nil || timeout < 0 {
		return utils.ToRESPError("invalid timeout")
	}

	// Get list keys (all args except the last one which is timeout)
	keys := args[:len(args)-1]

	// Check if any of the keys exist and have elements
	for _, key := range keys {
		existingObj, exists := c.Store.Caches[key]
		if !exists {
			continue
		}

		if existingObj.Type != store.LIST {
			return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
		}

		list, ok := existingObj.Value.([]string)
		if !ok {
			continue
		}

		if len(list) > 0 {
			// Found a non-empty list, pop from left
			popped := list[0]
			list = list[1:]

			// If list becomes empty, delete the key (Redis behavior)
			if len(list) == 0 {
				delete(c.Store.Caches, key)
			} else {
				c.Store.Caches[key] = store.CacheObject{
					Type:      store.LIST,
					Value:     list,
					TTL:       existingObj.TTL,
					CreatedAt: existingObj.CreatedAt,
				}
			}

			// Return array with key and value
			return utils.ToRESPArray([]string{key, popped})
		}
	}

	// If timeout is 0, return immediately
	if timeout == 0 {
		result, err := utils.EncodeRESP(nil)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	// For now, we'll implement a simple blocking mechanism
	// In a real implementation, this would use channels and goroutines
	// For simplicity, we'll just wait and then check again
	time.Sleep(time.Duration(timeout) * time.Second)

	// Check again after timeout
	for _, key := range keys {
		existingObj, exists := c.Store.Caches[key]
		if !exists {
			continue
		}

		if existingObj.Type != store.LIST {
			return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
		}

		list, ok := existingObj.Value.([]string)
		if !ok {
			continue
		}

		if len(list) > 0 {
			// Found a non-empty list, pop from left
			popped := list[0]
			list = list[1:]

			// If list becomes empty, delete the key (Redis behavior)
			if len(list) == 0 {
				delete(c.Store.Caches, key)
			} else {
				c.Store.Caches[key] = store.CacheObject{
					Type:      store.LIST,
					Value:     list,
					TTL:       existingObj.TTL,
					CreatedAt: existingObj.CreatedAt,
				}
			}

			// Return array with key and value
			return utils.ToRESPArray([]string{key, popped})
		}
	}

	// No elements found after timeout
	result, err := utils.EncodeRESP(nil)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
