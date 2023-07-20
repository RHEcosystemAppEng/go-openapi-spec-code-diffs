package positive

import (
	"github.com/RHEcosystemAppEng/openapi_spec_code_diffs/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeMatchesSpec(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, result := v.Validate()
	assert.Nil(t, err, "There is no error in execution: code matches specs")
	assert.Equal(t, 0, len(result.CodeDefsNotSpec), "There should be no mismatches: code should match specs")
}

func TestSpecMatchesCode(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, result := v.Validate()
	assert.Nil(t, err, "There is no error in execution: specs match code")
	assert.Equal(t, 0, len(result.SpecDefsNotInCode), "There should be no mismatches: specs should match code")
}
