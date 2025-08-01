package command

import (
	"time"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) DBSIZE(args []string) []byte {
	if len(args) != 0 {
		return utils.ToRESP("ERR wrong number of arguments for 'dbsize' command")
	}

	validKeys := 0
	for _, value := range c.Store.Caches {
		// Skip expired keys
		if value.TTL > 0 {
			elapsed := time.Now().Unix() - value.CreatedAt
			if value.TTL-time.Duration(elapsed) < 0 {
				continue
			}
		}
		validKeys++
	}

	result, err := utils.EncodeRESP(validKeys)
	if err != nil {
		return utils.ToRESP(err.Error())
	}
	return result
}
