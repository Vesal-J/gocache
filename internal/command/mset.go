package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) MSet(args []string) []byte {
	if len(args)%2 != 0 {
		return utils.ToRESPError("wrong number of arguments for 'mset' command")
	}

	for i := 0; i < len(args); i += 2 {
		key := args[i]
		value := args[i+1]

		c.Set([]string{key, value})
	}

	return utils.ToRESP("OK")
}
