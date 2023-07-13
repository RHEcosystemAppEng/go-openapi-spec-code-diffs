package main

import (
	"github.com/rs/zerolog/log"
)

// IgnoredAPISpecs Array of API paths to be ignored while working out diffs between code and specs
type IgnoredAPISpecs struct {
	ignoredAPISpecs []string
}

// NewIgnoredAPISpecs Returns a new IgnoredAPISpecs
func NewIgnoredAPISpecs(ignoredAPISpecsFile string) (error, *IgnoredAPISpecs) {
	err, ignoredSpecs := ReadFile(ignoredAPISpecsFile)
	if err != nil {
		log.Error().Msg("Error reading ignored specs file: " + ignoredAPISpecsFile + " " + err.Error())
		return err, nil
	}
	return nil, &IgnoredAPISpecs{ignoredAPISpecs: ignoredSpecs}
}
