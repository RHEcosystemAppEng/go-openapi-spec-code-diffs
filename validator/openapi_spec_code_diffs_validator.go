package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

// OpenAPISpecCodeDiffsValidator The main input struct for this validation tool
type OpenAPISpecCodeDiffsValidator struct {
	ignoredDirectories string
	ignoredAPISpecs    string
	goSourcesDir       string
	oasSpecFile        string // filename with path
}

// OpenAPISpecCodeDiffsResult Represents results of diff operation between golang source code and OpenAPI Specs
type OpenAPISpecCodeDiffsResult struct {
	SpecDefsNotInCode []string
	CodeDefsNotSpec   []string
}

// NewOpenAPISpecCodeDiffsValidator Returns a new OpenAPISpecCodeDiffsValidator
func NewOpenAPISpecCodeDiffsValidator(ignoredDirectories string, ignoredAPISpecs string, goSourcesDir string, oasSpecFile string) *OpenAPISpecCodeDiffsValidator {
	return &OpenAPISpecCodeDiffsValidator{ignoredDirectories: ignoredDirectories, ignoredAPISpecs: ignoredAPISpecs, goSourcesDir: goSourcesDir, oasSpecFile: oasSpecFile}
}

// setupValidator Sets up validator tool for its operation
func (v *OpenAPISpecCodeDiffsValidator) setupValidator() (error, *DiffWorker) {
	err, ignoredDirectories := NewIgnoredDirectories(v.ignoredDirectories)
	if err != nil {
		log.Error().Msg("Error while processing ignored directories " + err.Error())
		return err, nil
	}

	err, ignoredAPISpecs := NewIgnoredAPISpecs(v.ignoredAPISpecs)
	if err != nil {
		log.Error().Msg("Error while processing ignored specs" + err.Error())
		return err, nil
	}

	dirWalker := NewGoSourceDirectory(v.goSourcesDir, ignoredDirectories)
	err, codeAPIDefs := dirWalker.WalkDirectory()
	if err != nil {
		log.Error().Msg("Error while processing go source directories" + err.Error())
		return err, nil
	}

	oasModel := NewOASModel(v.oasSpecFile)
	oasModel.LoadSpecModel()
	specAPIDefs := oasModel.GetPathOps()

	return nil, NewDiffWorker(codeAPIDefs, specAPIDefs, ignoredAPISpecs)
}

// Validate main validation function that sets up the validation tool and returns the diff result
func (v *OpenAPISpecCodeDiffsValidator) Validate() (error, OpenAPISpecCodeDiffsResult) {
	err, validator := v.setupValidator()
	if err != nil {
		log.Error().Msg("Error setting up new validator: " + err.Error())
		return err, OpenAPISpecCodeDiffsResult{nil, nil}
	}

	err, codeDefsNotInSpec := validator.ValidateCodeDefsNotInSpec()
	if err != nil {
		log.Error().Msg("Error while processing code defs not in spec diffs " + err.Error())
		return err, OpenAPISpecCodeDiffsResult{nil, nil}
	}
	if len(codeDefsNotInSpec) == 0 {
		log.Info().Msg("Successful validation 1: endpoints defined in code are present in api specs")
	} else {
		log.Error().Msg(fmt.Sprintf("Unsuccessful validation 1: %d defs in code not defined in spec.", len(codeDefsNotInSpec)))
	}

	err, specDefsNotInCode := validator.ValidateSpecDefsNotInCode()
	if err != nil {
		log.Error().Msg("Error while processing spec defs not in code diffs " + err.Error())
		return err, OpenAPISpecCodeDiffsResult{nil, nil}
	}
	if len(specDefsNotInCode) == 0 {
		log.Info().Msg("Successful validation 2: endpoints defined in spec are present in code")
	} else {
		log.Error().Msg(fmt.Sprintf("Unsuccessful validation 2: %d defs in spec not implemented in code.", len(specDefsNotInCode)))
	}

	return nil, OpenAPISpecCodeDiffsResult{specDefsNotInCode, codeDefsNotInSpec}
}
