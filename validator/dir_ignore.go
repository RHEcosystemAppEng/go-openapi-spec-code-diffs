package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

// IgnoredDirectories Directories to be ignored while coming up with diffs e.g. /bin, /out
type IgnoredDirectories struct {
	ignoredDirectories []string
}

// NewIgnoredDirectories Returns new IgnoredDirectories
func NewIgnoredDirectories(ignoreDirFileName string) (error, *IgnoredDirectories) {
	dirs := []string{}
	if strings.TrimSpace(ignoreDirFileName) == "" {
		log.Debug().Msg(fmt.Sprintf("No ignored directory file name specified, using empty ignored dirs"))
		return nil, &IgnoredDirectories{ignoredDirectories: dirs}
	}

	err, dirs := ReadFile(ignoreDirFileName)
	if err != nil {
		log.Error().Msg("Error reading ignored directory file: " + ignoreDirFileName + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredDirectories{ignoredDirectories: dirs}
}
