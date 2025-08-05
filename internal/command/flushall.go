package command

import (
	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) FlushAll(args []string) []byte {
	if len(args) != 0 {
		return utils.ToRESPError("wrong number of arguments for 'flushall' command")
	}

	c.Store.Caches = make(map[string]store.CacheObject)

	return utils.ToRESP("OK")
}
