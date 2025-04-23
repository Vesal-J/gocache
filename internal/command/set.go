package command

import (
	"fmt"
	"strconv"

	"github.com/vesal-j/gocache/internal/store"
)

// SET key value TTL
// this is the order of args, here in this function we are getting key, value and ttl by this order
func (c *CommandImpl) Set(args []string) string {
	if len(args) < 2 {
		return "ERR wrong number of arguments for 'set' command"
	}

	TtlNumber := 0
	if len(args) == 3 {
		var err error
		TtlNumber, err = strconv.Atoi(args[2])
		if err != nil {
			return "ERR invalid TTL"
		}
	}

	c.Store.Caches[args[0]] = store.CacheObject{
		Key:   args[0],
		Value: args[1],
		TTL:   TtlNumber,
	}
	fmt.Println(c.Store.Caches[args[0]])
	return "OK"
}
