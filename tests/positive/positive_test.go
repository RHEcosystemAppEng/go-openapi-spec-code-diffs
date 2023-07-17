package positive

import (
	"github.com/RHEcosystemAppEng/openapi_spec_code_diffs/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeMatchesSpec(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, result := v.Validate()
	assert.Nil(t, err, "There is no error in execution: code matches specs")
	assert.Equal(t, 0, len(result.CodeDefsNotSpec), "There are no mismatches: code matches specs")
}

func TestSpecMatchesCode(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, result := v.Validate()
	assert.Nil(t, err, "There is no error in execution: specs match code")
	assert.Equal(t, 0, len(result.SpecDefsNotInCode), "There are no mismatches: specs match code")
}
