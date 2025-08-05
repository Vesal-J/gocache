package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) RenameNX(args []string) []byte {
	if len(args) != 2 {
		return utils.ToRESPError("wrong number of arguments for 'renamenx' command")
	}

	oldKey := args[0]
	newKey := args[1]

	if _, exists := c.Store.Caches[oldKey]; !exists {
		return utils.ToRESPError("ERR no such key")
	}

	if _, exists := c.Store.Caches[newKey]; exists {
		return utils.ToRESP("0")
	}

	c.Store.Caches[newKey] = c.Store.Caches[oldKey]
	delete(c.Store.Caches, oldKey)

	return utils.ToRESP("1")
}
