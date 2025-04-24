package command

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Scan(args []string) []byte {
	if len(args) < 1 {
		return utils.ToRESP("ERR wrong number of arguments for 'scan' command")
	}

	_, err := strconv.Atoi(args[0])
	if err != nil {
		return utils.ToRESP("ERR invalid cursor")
	}

	var pattern string = "*"
	var count int = 10 // default count

	// Parse optional arguments
	for i := 1; i < len(args); i++ {
		switch strings.ToUpper(args[i]) {
		case "MATCH":
			if i+1 >= len(args) {
				return utils.ToRESP("ERR syntax error")
			}
			pattern = args[i+1]
			i++
		case "COUNT":
			if i+1 >= len(args) {
				return utils.ToRESP("ERR syntax error")
			}
			count, err = strconv.Atoi(args[i+1])
			if err != nil {
				return utils.ToRESP("ERR value is not an integer or out of range")
			}
			i++
		}
	}

	// Get all keys that match the pattern
	matches := make([]string, 0)
	for key := range c.Store.Caches {
		matched, err := filepath.Match(pattern, key)
		if err != nil {
			continue // Skip invalid patterns
		}
		if matched {
			matches = append(matches, key)
			if len(matches) >= count {
				break
			}
		}
	}

	// Return matches with cursor
	result, err := utils.EncodeRESP([]any{"0", matches})
	if err != nil {
		return utils.ToRESP(err.Error())
	}

	return result
}
