package command

import (
	"strconv"
	"strings"
	"time"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Set(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'set' command")
	}

	key := args[0]
	value := args[1]
	ttl := 0
	nx := false
	xx := false

	i := 2
	for i < len(args) {
		arg := strings.ToUpper(args[i])

		switch arg {
		case "NX":
			nx = true
			i++
		case "XX":
			xx = true
			i++
		case "EX":
			if i+1 >= len(args) {
				return utils.ToRESPError("syntax error")
			}
			seconds, err := strconv.Atoi(args[i+1])
			if err != nil || seconds < 0 {
				return utils.ToRESPError("invalid TTL")
			}
			ttl = seconds
			i += 2
		case "PX":
			if i+1 >= len(args) {
				return utils.ToRESPError("syntax error")
			}
			millis, err := strconv.Atoi(args[i+1])
			if err != nil || millis < 0 {
				return utils.ToRESPError("invalid TTL")
			}
			ttl = millis / 1000
			i += 2
		default:
			return utils.ToRESPError("syntax error")
		}
	}

	_, exists := c.Store.Caches[key]

	if nx && exists {
		result, err := utils.EncodeRESP(nil)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result // do not set if exists
	}
	if xx && !exists {
		result, err := utils.EncodeRESP(nil)
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result // do not set if not exists
	}

	c.Store.Caches[key] = store.CacheObject{
		Key:       key,
		Value:     value,
		TTL:       ttl,
		CreatedAt: time.Now().Unix(),
	}

	return utils.ToRESP("OK")
}
