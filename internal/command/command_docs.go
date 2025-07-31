package command

import (
	"github.com/vesal-j/gocache/internal/utils"
)

func (c *CommandImpl) Command(args []string) []byte {
	if len(args) == 0 {
		// Return empty array for COMMAND without arguments
		result, err := utils.EncodeRESP([]string{})
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	if len(args) == 1 && args[0] == "DOCS" {
		// Return empty array for COMMAND DOCS
		result, err := utils.EncodeRESP([]string{})
		if err != nil {
			return utils.ToRESPError(err.Error())
		}
		return result
	}

	// For any other COMMAND subcommand, return empty array
	result, err := utils.EncodeRESP([]string{})
	if err != nil {
		return utils.ToRESPError(err.Error())
	}
	return result
}
