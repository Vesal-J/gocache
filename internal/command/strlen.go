package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Strlen(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'strlen' command")
	}

	key := args[0]
	value, exists := c.Store.Caches[key]
	if !exists {
		result, err := utils.EncodeRESP(0)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	result, err := utils.EncodeRESP(len(value.Value))
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
