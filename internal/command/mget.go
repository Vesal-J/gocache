package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) MGet(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESPError("wrong number of arguments for 'mget' command")
	}

	// Create an array to hold the results
	var results []any

	for _, key := range args {
		value, exists := c.Store.Caches[key]
		if !exists {
			results = append(results, nil)
		} else {
			results = append(results, value.Value)
		}
	}

	encoded, err := utils.EncodeRESP(results)
	if err != nil {
		return utils.ToRESPError("internal error encoding response")
	}

	return encoded
}
