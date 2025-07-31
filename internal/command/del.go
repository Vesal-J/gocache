package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Del(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESPError("wrong number of arguments for 'del' command")
	}

	deletedCount := 0

	for _, arg := range args {
		_, exists := c.Store.Caches[arg]
		if exists {
			delete(c.Store.Caches, arg)
			deletedCount++
		}
	}

	result, err := utils.EncodeRESP(deletedCount)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
