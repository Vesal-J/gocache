package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Ping(args []string) []byte {
	if len(args) == 0 {
		return utils.ToRESP("PONG")
	}

	return []byte(args[0])
}
