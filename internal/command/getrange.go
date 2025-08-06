package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Getrange(args []string) []byte {
	if len(args) != 3 {
		return utils.ToRESPError("wrong number of arguments for 'getrange' command")
	}

	key := args[0]
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("value is not an integer or out of range")
	}

	end, err := strconv.Atoi(args[2])
	if err != nil {
		return utils.ToRESPError("value is not an integer or out of range")
	}

	value, exists := c.Store.Caches[key]
	if !exists {
		result, err := utils.EncodeRESP(nil)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	// Handle negative indices
	strLen := len(value.Value.(string))
	if start < 0 {
		start = strLen + start
	}
	if end < 0 {
		end = strLen + end
	}

	// Clamp indices to string bounds
	if start < 0 {
		start = 0
	}
	if end >= strLen {
		end = strLen - 1
	}
	if start > end {
		return utils.ToRESP("")
	}

	// Extract substring
	substring := value.Value.(string)[start : end+1]
	return utils.ToRESP(substring)
}
