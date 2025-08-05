package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Rename(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'rename' command")
	}

	oldKey := args[0]
	newKey := args[1]

	if _, exists := c.Store.Caches[oldKey]; !exists {
		return utils.ToRESPError("ERR no such key")
	}

	c.Store.Caches[newKey] = c.Store.Caches[oldKey]
	delete(c.Store.Caches, oldKey)

	return utils.ToRESP("OK")
}
