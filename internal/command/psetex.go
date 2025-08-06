package command

import (
	"strconv"
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) PSetEX(args []string) []byte {
	if len(args) < 3 {
		return utils.ToRESPError("wrong number of arguments for 'psetex' command")
	}

	ttl, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("invalid ttl")
	}

	c.Store.Caches[args[0]] = store.CacheObject{
		Value:     args[2],
		TTL:       time.Duration(ttl) * time.Millisecond,
		CreatedAt: time.Now().Unix(),
	}

	return utils.ToRESP("OK")
}
