package command

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/vesal-j/gocache/internal/store"
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) HScan(args []string) []byte {
	if len(args) < 2 {
		return utils.ToRESPError("wrong number of arguments for 'hscan' command")
	}

	key := args[0]
	_, err := strconv.Atoi(args[1])
	if err != nil {
		return utils.ToRESPError("ERR invalid cursor")
	}

	cacheObj, exists := c.Store.Caches[key]
	if !exists {
		result, err := utils.EncodeRESP([]any{"0", []string{}})
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	if cacheObj.Type != store.HASH {
		return utils.ToRESPError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	hash := cacheObj.Value.(map[string]string)
	var pattern string = "*"
	var count int = 10

	for i := 2; i < len(args); i++ {
		switch strings.ToUpper(args[i]) {
		case "MATCH":
			if i+1 >= len(args) {
				return utils.ToRESPError("ERR syntax error")
			}
			pattern = args[i+1]
			i++
		case "COUNT":
			if i+1 >= len(args) {
				return utils.ToRESPError("ERR syntax error")
			}
			count, err = strconv.Atoi(args[i+1])
			if err != nil {
				return utils.ToRESPError("ERR value is not an integer or out of range")
			}
			i++
		}
	}

	matches := make([]string, 0)
	for field, value := range hash {
		matched, err := filepath.Match(pattern, field)
		if err != nil {
			continue
		}
		if matched {
			matches = append(matches, field, value)
			if len(matches) >= count*2 {
				break
			}
		}
	}

	result, err := utils.EncodeRESP([]any{"0", matches})
	if err != nil {
		return utils.ToRESPError(err.Error())
	}

	return result
}
