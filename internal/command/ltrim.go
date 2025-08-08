package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) LTrim(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'ltrim' command")
	}

	key := args[0]
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid start index")
	}

	stop, err := strconv.Atoi(args[2])
	if err != nil {
		return utils.ToRESPError("invalid stop index")
	}

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

	// Handle negative indices (Redis behavior)
	if start < 0 {
		start = len(list) + start
	}
	if start < 0 {
		start = 0
	}

	if stop < 0 {
		stop = len(list) + stop
	}
	if stop >= len(list) {
		stop = len(list) - 1
	}

	// If start is beyond the list length or start > stop, trim to empty list
	if start >= len(list) || start > stop {
		// Delete the key if trimming to empty (Redis behavior)
		delete(c.Store.Caches, key)
		return utils.ToRESP("OK")
	}

	// Ensure stop doesn't exceed list bounds
	if stop >= len(list) {
		stop = len(list) - 1
	}

	// Trim the list to the specified range
	trimmedList := list[start : stop+1]

	// Update the cache object
	c.Store.Caches[key] = store.CacheObject{
		Type:      store.LIST,
		Value:     trimmedList,
		TTL:       existingObj.TTL,
		CreatedAt: existingObj.CreatedAt,
	}

	return utils.ToRESP("OK")
}
