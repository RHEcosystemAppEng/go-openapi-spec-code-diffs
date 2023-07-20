package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

// IgnoredFiles Files to be ignored while coming up with diffs
type IgnoredFiles struct {
	ignoredFiles []string
}

// NewIgnoredFiles Returns new IgnoredFiles
func NewIgnoredFiles(ignoreFilesFileName string) (error, *IgnoredFiles) {
	files := []string{}
	if strings.TrimSpace(ignoreFilesFileName) == "" {
		log.Debug().Msg(fmt.Sprintf("No ignored files file name specified, using empty ignored files"))
		return nil, &IgnoredFiles{ignoredFiles: files}
	}

	err, files := ReadFile(ignoreFilesFileName)
	if err != nil {
		log.Error().Msg("Error reading ignored files file: " + ignoreFilesFileName + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredFiles{ignoredFiles: files}
}
