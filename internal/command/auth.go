package command

import "github.com/vesal-j/gocache/internal/utils"

func (c *CommandImpl) Auth(args []string) []byte {
	return utils.ToRESP("OK")
}
