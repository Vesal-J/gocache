package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) GetSet(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'getset' command")
	}

	key := args[0]
	value := args[1]

	oldValue, exists := c.Store.Caches[key]

	if exists {
		c.Set([]string{key, value})
		return utils.ToRESP(oldValue.Value)
	} else {
		c.Set([]string{key, value})
		return utils.ToRESP("(nil)")
	}
}
