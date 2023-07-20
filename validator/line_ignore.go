package validator

import "github.com/rs/zerolog/log"

// IgnoredLines Files to be ignored while coming up with diffs
type IgnoredLines struct {
	ignoredLines []string
}

// NewIgnoredLines Returns new IgnoredLines
func NewIgnoredLines(ignoreFilesFileName string) (error, *IgnoredLines) {
	err, lines := ReadFile(ignoreFilesFileName)
	if err != nil {
		log.Error().Msg("Error reading ignored files file: " + ignoreFilesFileName + " " + err.Error())
		return err, nil
	}

	return nil, &IgnoredLines{ignoredLines: lines}
}
