package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Type(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'type' command")
	}

	key := args[0]
	_, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("none")
	}

	// Currently we only support string type
	return utils.ToRESP("string")
}
