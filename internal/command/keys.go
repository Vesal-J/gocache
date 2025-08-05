package command

import (
	"path/filepath"
	"strings"

	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Keys(args []string) []byte {
	if len(args) != 1 {
		return utils.ToRESPError("wrong number of arguments for 'keys' command")
	}

	pattern := args[0]
	var matchingKeys []string

	for key := range c.Store.Caches {
		if matchPattern(key, pattern) {
			matchingKeys = append(matchingKeys, key)
		}
	}

	return utils.ToRESPArray(matchingKeys)
}

func matchPattern(key, pattern string) bool {
	filePattern := strings.ReplaceAll(pattern, "*", "*")
	filePattern = strings.ReplaceAll(filePattern, "?", "?")

	matched, err := filepath.Match(filePattern, key)
	if err != nil {
		return false
	}

	return matched
}
