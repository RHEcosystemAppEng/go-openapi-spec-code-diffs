package validator

import (
	"github.com/rs/zerolog/log"
)

// IgnoredAPIPaths Array of API paths to be ignored while working out diffs between code and specs
type IgnoredAPIPaths struct {
	ignoredAPIPaths []string
}

// NewIgnoredAPIRoutes Returns a new IgnoredAPIPaths
func NewIgnoredAPIRoutes(ignoredAPIPathsFile string) (error, *IgnoredAPIPaths) {
	err, ignoredPaths := ReadFile(ignoredAPIPathsFile)
	if err != nil {
		log.Error().Msg("Error reading ignored api paths file: " + ignoredAPIPathsFile + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredAPIPaths{ignoredAPIPaths: ignoredPaths}
}
