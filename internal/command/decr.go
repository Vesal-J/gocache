package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Decr(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESPError("wrong number of arguments for 'decr' command")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		c.Set([]string{args[0], "-1"})
		return utils.ToRESP("-1")
	}

	num, err := strconv.Atoi(value.Value)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}

	num--
	c.Set([]string{args[0], strconv.Itoa(num)})

	return utils.ToRESP(strconv.Itoa(num))
}
