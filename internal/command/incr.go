package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Incr(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESPError("wrong number of arguments for 'incr' command")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		c.Set([]string{args[0], "1"})
		return utils.ToRESP("1")
	}

	num, err := strconv.Atoi(value.Value.(string))
	if err != nil {
		return utils.ToRESPError(err.Error())
	}

	num++
	c.Set([]string{args[0], strconv.Itoa(num)})

	return utils.ToRESP(strconv.Itoa(num))
}
