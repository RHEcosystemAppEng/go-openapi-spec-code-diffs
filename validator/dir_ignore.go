package validator

import "github.com/rs/zerolog/log"

// IgnoredDirectories Directories to be ignored while coming up with diffs e.g. /bin, /out
type IgnoredDirectories struct {
	ignoredDirectories []string
}

// NewIgnoredDirectories Returns new IgnoredDirectories
func NewIgnoredDirectories(ignoreDirFileName string) (error, *IgnoredDirectories) {
	err, dirs := ReadFile(ignoreDirFileName)
	if err != nil {
		log.Error().Msg("Error reading ignored directory file: " + ignoreDirFileName + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredDirectories{ignoredDirectories: dirs}
}
