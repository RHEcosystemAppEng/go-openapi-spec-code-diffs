package main

import (
	"flag"
	"github.com/RHEcosystemAppEng/openapi_spec_code_diffs/validator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

const DIFFS_FOUND = -1
const SUCCESS = 0
const ERROR_PROCESSING_DIFFS = -2
const ERROR_INSUFFICIENT_ARGS = -3

var ignoredDirsFile *string
var ignoredFilesFile *string
var ignoredLinesFile *string
var ignoredPathsFile *string
var goSourceDir *string
var openAPISpecsFile *string
var logLevelArg *string
var logLevel zerolog.Level = zerolog.InfoLevel

func main() {
	ignoredDirsFile = flag.String("ignoredDirsFile", "", "File containing directories that are ignored")
	ignoredFilesFile = flag.String("ignoredFilesFile", "", "File containing file names that are ignored")
	ignoredLinesFile = flag.String("ignoredLinesFile", "", "File containing lines that are ignored")
	ignoredPathsFile = flag.String("ignoredPathsFile", "", "File containing api paths that are ignored")
	goSourceDir = flag.String("goSourceDir", "./", "Path to go sources directory")
	openAPISpecsFile = flag.String("openAPISpecsFile", "./openapi.yaml", "Filename including path to openapi specifications")
	logLevelArg = flag.String("logLevel", "info", "Log level: disabled, info, debug, error")
	flag.Parse()
	mapLogLevelArg()
	zerolog.SetGlobalLevel(logLevel)

	validator := validator.NewOpenAPISpecCodeDiffsValidator(*ignoredDirsFile, *ignoredFilesFile, *ignoredLinesFile, *ignoredPathsFile, *goSourceDir, *openAPISpecsFile)
	err, result := validator.Validate()
	if err != nil {
		log.Error().Msg("Error while performing diffs: " + err.Error())
		os.Exit(ERROR_PROCESSING_DIFFS)
	}

	diffsFound := len(result.SpecDefsNotInCode) > 0 || len(result.CodeDefsNotSpec) > 0

	if result.SpecDefsNotInCode != nil && len(result.SpecDefsNotInCode) > 0 {
		log.Warn().Msg("API Specs not found in code: " + strconv.Itoa(len(result.SpecDefsNotInCode)))
	}

	if result.CodeDefsNotSpec != nil && len(result.CodeDefsNotSpec) > 0 {
		log.Warn().Msg("Code Defs not found in API specs: " + strconv.Itoa(len(result.CodeDefsNotSpec)))
	}

	if diffsFound {
		log.Error().Msg("Unsuccessful: Diffs found")
		os.Exit(DIFFS_FOUND)
	} else {
		log.Info().Msg("Success: No diffs found")
		os.Exit(SUCCESS)
	}
}

// mapLogLevelArg Maps program arguments to variables
func mapLogLevelArg() {
	switch *logLevelArg {
	case "disabled":
		logLevel = zerolog.Disabled
	case "info":
		logLevel = zerolog.InfoLevel
	case "debug":
		logLevel = zerolog.DebugLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		log.Error().Msg("Invalid log level passed, will default to Info")
	}
}
