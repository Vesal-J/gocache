package command

import (
	"strings"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Client(args []string) []byte {
	if len(args) == 0 {
		return utils.ToRESP("ERR wrong number of arguments for 'client' command")
	}

	subcommand := strings.ToUpper(args[0])
	switch subcommand {
	case "LIST":
		return utils.ToRESP("id=1 addr=127.0.0.1:6379 fd=5 name= age=3 idle=0 flags=N db=0 sub=0 psub=0 multi=-1 qbuf=0 qbuf-free=0 obl=0 oll=0 omem=0 events=r cmd=ping")
	case "SETNAME":
		if len(args) != 2 {
			return utils.ToRESP("ERR wrong number of arguments for 'client setname' command")
		}
		return utils.ToRESP("OK")
	default:
		return utils.ToRESP("ERR unknown subcommand '" + subcommand + "' for 'client' command")
	}
}
