package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
)

// DiffWorker Structure for use as inputs by diff worker, consisting of definitions found in code, specs and api definitions to be ignored
type DiffWorker struct {
	codeAPIDefs     []*goSourceAPILine
	specAPIDefs     []string
	ignoredAPISpecs *IgnoredAPIPaths
}

// NewDiffWorker Returns a new DiffWorker
func NewDiffWorker(codeAPIDefs []*goSourceAPILine, specAPIDefs []string, ignoredAPIDefs *IgnoredAPIPaths) *DiffWorker {
	return &DiffWorker{
		codeAPIDefs: codeAPIDefs, specAPIDefs: specAPIDefs, ignoredAPISpecs: ignoredAPIDefs,
	}
}

// ValidateCodeDefsNotInSpec Returns api paths found in code but not defined in spec
func (v *DiffWorker) ValidateCodeDefsNotInSpec() (error, []string) {
	var diffs []string
	foundInSpec := make(map[string]string)
	for _, apiDef := range v.specAPIDefs {
		foundInSpec[apiDef] = apiDef
	}

	for _, apiDef := range v.codeAPIDefs {
		if !slices.Contains(v.ignoredAPISpecs.ignoredAPIPaths, apiDef.apiPath) {
			if len(apiDef.httpMethod) > 0 {
				v := foundInSpec[apiDef.httpMethod+" "+apiDef.apiPath]
				if len(v) == 0 {
					log.Info().Msg(fmt.Sprintf("API Def at line %d in file '%s' with '%s %s' not found in spec:", apiDef.lineNum, apiDef.fileName, apiDef.httpMethod, apiDef.apiPath))
					diffs = append(diffs, apiDef.apiPath)
				}
			}
		}
	}

	return nil, diffs
}

// ValidateSpecDefsNotInCode Returns api paths defined in spec but not found in code
func (v *DiffWorker) ValidateSpecDefsNotInCode() (error, []string) {
	var diffs []string
	foundInCode := make(map[string]*goSourceAPILine)
	for _, apiDef := range v.codeAPIDefs {
		if len(apiDef.httpMethod) > 0 {
			key := apiDef.httpMethod + " " + apiDef.apiPath
			foundInCode[key] = apiDef
		}
	}

	for _, apiDef := range v.specAPIDefs {
		if !slices.Contains(v.ignoredAPISpecs.ignoredAPIPaths, apiDef) {
			v := foundInCode[apiDef]
			if v == nil {
				log.Info().Msg(fmt.Sprintf("API defined in spec does not have corresponding implementation %s:", apiDef))
				diffs = append(diffs, apiDef)
			}
		}
	}

	return nil, diffs
}
