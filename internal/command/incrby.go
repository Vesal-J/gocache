package command

import (
	"strconv"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) IncrBy(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'incrby' command")
	}

	increment, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid increment")
	}

	value, exists := c.Store.Caches[args[0]]
	if !exists {
		c.Set([]string{args[0], strconv.Itoa(increment)})
		return utils.ToRESP(strconv.Itoa(increment))
	}

	num, err := strconv.Atoi(value.Value)
	if err != nil {
		return utils.ToRESPError(err.Error())
	}

	num += increment
	c.Set([]string{args[0], strconv.Itoa(num)})

	return utils.ToRESP(strconv.Itoa(num))
}
