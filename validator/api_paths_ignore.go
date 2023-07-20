package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

// IgnoredAPIPaths Array of API paths to be ignored while working out diffs between code and specs
type IgnoredAPIPaths struct {
	ignoredAPIPaths []string
}

// NewIgnoredAPIRoutes Returns a new IgnoredAPIPaths
func NewIgnoredAPIRoutes(ignoredAPIPathsFile string) (error, *IgnoredAPIPaths) {
	paths := []string{}
	if strings.TrimSpace(ignoredAPIPathsFile) == "" {
		log.Debug().Msg(fmt.Sprintf("No ignored api paths file name specified, using empty ignored api paths"))
		return nil, &IgnoredAPIPaths{ignoredAPIPaths: paths}
	}

	err, ignoredPaths := ReadFile(ignoredAPIPathsFile)
	if err != nil {
		log.Error().Msg("Error reading ignored api paths file: " + ignoredAPIPathsFile + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredAPIPaths{ignoredAPIPaths: ignoredPaths}
}
