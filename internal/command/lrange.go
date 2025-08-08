package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) LRange(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'lrange' command")
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

	if start < 0 {
		start = len(list) + start
	}

	if stop < 0 {
		stop = len(list) + stop
	}

	if start > stop {
		return utils.ToRESPArray([]string{})
	}

	rangeList := list[start : stop+1]

	return utils.ToRESPArray(rangeList)
}
