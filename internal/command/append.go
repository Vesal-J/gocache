package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Append(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'append' command")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		c.Set([]string{args[0], args[1]})
		return utils.ToRESP(strconv.Itoa(len(args[1])))
	} else {
		c.Set([]string{args[0], value.Value + args[1]})
		return utils.ToRESP(strconv.Itoa(len(value.Value) + len(args[1])))
	}
}
