package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Exists(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'exists' command")
	}

	_, exists := c.Store.Caches[args[0]]
	if !exists {
		result, err := utils.EncodeRESP(0)
		if err != nil {
			return utils.ToRESP(err.Error())
		}
		return result
	}

	result, err := utils.EncodeRESP(1)
	if err != nil {
		return utils.ToRESP(err.Error())
	}
	return result
}
