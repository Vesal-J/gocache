package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Get(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'get' command")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		return utils.ToRESP("(nil)")
	}

	return utils.ToRESP(value.Value)
}
