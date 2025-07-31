package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Get(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'get' command")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		result, err := utils.EncodeRESP(nil)
		if err != nil {
			return utils.ToRESP(err.Error())
		}
		return result
	}

	return utils.ToRESP(value.Value)
}
