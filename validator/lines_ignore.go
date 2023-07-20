package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

// IgnoredLines Files to be ignored while coming up with diffs
type IgnoredLines struct {
	ignoredLines []string
}

// NewIgnoredLines Returns new IgnoredLines
func NewIgnoredLines(ignoredLinesFileName string) (error, *IgnoredLines) {
	lines := []string{}
	if strings.TrimSpace(ignoredLinesFileName) == "" {
		log.Debug().Msg(fmt.Sprintf("No ignored lines file name specified, using empty ignored files"))
		return nil, &IgnoredLines{ignoredLines: lines}
	}

	err, lines := ReadFile(ignoredLinesFileName)
	if err != nil {
		log.Error().Msg("Error reading ignored files file: " + ignoredLinesFileName + " " + err.Error())
		return err, nil
	}

	return nil, &IgnoredLines{ignoredLines: lines}
}
