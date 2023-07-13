package main

import (
	"fmt"
	"github.com/RHEcosystemAppEng/openapi_spec_code_diffs/openapi-spec-code-diffs"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

const DIFFS_FOUND = -1
const SUCCESS = 0
const ERROR_PROCESSING_DIFFS = -2
const ERROR_INSUFFICIENT_ARGS = -3

var ignoredDirsFile string
var ignoredPathsFile string
var goSourceDir string
var openAPIFile string

func main() {
	if len(os.Args) < 5 {
		showUsage()
		os.Exit(ERROR_INSUFFICIENT_ARGS)
	}
	argsToVars()
	validator := openapi_spec_code_diffs.NewOpenAPISpecCodeDiffsValidator(ignoredDirsFile, ignoredPathsFile, goSourceDir, openAPIFile)
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

// argsToVars Maps program arguments to variables
func argsToVars() {
	openAPIFile = os.Args[1]
	goSourceDir = os.Args[2]
	ignoredDirsFile = os.Args[3]
	ignoredPathsFile = os.Args[4]
}

// showUsage Shows program usage
func showUsage() {
	fmt.Println("Usage: openapi-spec-code-diffs 'path/to/openapi/specs/filename' 'path/to/golang/source/dir' 'path/to/ignored/directories/filename' 'path/to/ignored/paths/filename'")
	fmt.Println("Usage Example: openapi-spec-code-diffs '~/example-service/openapi.yaml' '~/example-service' '~/example-service/.dirignore' '~/example-service/.specignore'")
}
