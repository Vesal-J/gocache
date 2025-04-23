package command

func (c *CommandImpl) Get(args []string) string {
	if len(args) != 1 {
		return "ERR wrong number of arguments for 'get' command"
	}

	// Check if key exists in cache
	value, exists := c.Store.Caches[args[0]]
	if !exists {
		return "(nil)"
	}

	return value.Value
}
