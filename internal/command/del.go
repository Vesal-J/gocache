package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Del(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESPError("wrong number of arguments for 'del' command")
	}

	_, exists := c.Store.Caches[args[0]]
	if !exists {
		result, err := utils.EncodeRESP(0)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	delete(c.Store.Caches, args[0])

	result, err := utils.EncodeRESP(1)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
