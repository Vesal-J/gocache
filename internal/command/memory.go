package command

import (
	"strings"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Memory(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'memory' command")
	}

	subcommand := strings.ToUpper(args[0])
	switch subcommand {
	case "USAGE":
		if len(args) < 2 {
			return utils.ToRESP("ERR wrong number of arguments for 'memory usage' command")
		}
		key := args[1]
		value, exists := c.Store.Caches[key]
		if !exists {
			return utils.ToRESP("(nil)")
		}

		// Calculate approximate memory usage
		// Key size + Value size + struct overhead
		memoryUsage := len(value.Key) + len(value.Value) + 24 // 24 bytes for struct overhead
		result, err := utils.EncodeRESP(memoryUsage)
		if err != nil {
			return utils.ToRESP(err.Error())
		}
		return result

	case "SAMPLES":
		// For now, just return 0 as we don't implement sampling yet
		result, err := utils.EncodeRESP(0)
		if err != nil {
			return utils.ToRESP(err.Error())
		}
		return result

	default:
		return utils.ToRESP("ERR unknown subcommand for 'memory' command")
	}
}
