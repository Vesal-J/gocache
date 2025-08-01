package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) DecrBy(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'decrby' command")
	}

	decrement, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid decrement")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		c.Set([]string{args[0], strconv.Itoa(-decrement)})
		return utils.ToRESP(strconv.Itoa(-decrement))
	}

	num, err := strconv.Atoi(value.Value)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}

	num -= decrement
	c.Set([]string{args[0], strconv.Itoa(num)})

	return utils.ToRESP(strconv.Itoa(num))
}
