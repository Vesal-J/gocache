package command

import (
	"fmt"
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
			_, err := strconv.Atoi(args[i+1])
			if err != nil {
				return utils.ToRESP("ERR value is not an integer or out of range")
			}
			i++
		}
	}

	// Get all keys that match the pattern
	matches := make([]string, 0)
	for key := range c.Store.Caches {
		if strings.Contains(key, pattern) {
			matches = append(matches, key)
		}
	}

	// Return all matches since we don't implement actual cursor-based iteration
	result, err := utils.EncodeRESP([]any{"0", matches}) // matches is []string
	if err != nil {
		return utils.ToRESP(err.Error())
	}

	fmt.Println(c.Store.Caches)
	return result
}
