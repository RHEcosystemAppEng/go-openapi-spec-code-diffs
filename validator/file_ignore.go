package validator

import "github.com/rs/zerolog/log"

// IgnoredFiles Files to be ignored while coming up with diffs
type IgnoredFiles struct {
	ignoredFiles []string
}

// NewIgnoredFiles Returns new IgnoredFiles
func NewIgnoredFiles(ignoreFilesFileName string) (error, *IgnoredFiles) {
	err, files := ReadFile(ignoreFilesFileName)
	if err != nil {
		log.Error().Msg("Error reading ignored files file: " + ignoreFilesFileName + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredFiles{ignoredFiles: files}
}
