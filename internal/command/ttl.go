package command

import (
	"time"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) TTL(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'ttl' command")
	}

	key := args[0]
	value, exists := c.Store.Caches[key]
	if !exists {
		return utils.ToRESP("-2") // Key does not exist
	}

	// If TTL is 0, key is persistent
	if value.TTL == 0 {
		return utils.ToRESP("-1") // Key exists but has no TTL
	}

	// Calculate remaining TTL
	elapsed := time.Now().Unix() - value.CreatedAt
	remaining := value.TTL - int(elapsed)

	// If remaining time is negative, key has expired
	if remaining < 0 {
		return utils.ToRESP("-2") // Key has expired
	}

	result, err := utils.EncodeRESP(remaining)
	if err != nil {
		return utils.ToRESP(err.Error())
	}
	return result
}
