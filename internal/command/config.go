package command

import (
	"strings"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Config(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'config' command")
	}

	subcommand := strings.ToUpper(args[0])
	switch subcommand {
	case "GET":
		if len(args) != 2 {
			return utils.ToRESP("ERR wrong number of arguments for 'config get' command")
		}
		param := strings.ToLower(args[1])
		switch param {
		case "databases":
			result, err := utils.EncodeRESP([]string{"databases", "16"})
			if err != nil {
				return utils.ToRESP(err.Error())
			}
			return result
		default:
			return utils.ToRESP("ERR unknown subcommand '" + subcommand + "' for 'config' command")
		}
	default:
		return utils.ToRESP("ERR unknown subcommand '" + subcommand + "' for 'config' command")
	}
}
