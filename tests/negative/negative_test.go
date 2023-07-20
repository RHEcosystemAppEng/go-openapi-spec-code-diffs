package positive

import (
	"github.com/RHEcosystemAppEng/openapi_spec_code_diffs/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeDoesNotMatchSpec(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, result := v.Validate()
	assert.Nil(t, err, "There is no error in execution: code matches specs")
	assert.Equal(t, 1, len(result.CodeDefsNotSpec), "There is exactly 1 mismatch: code does not match specs")
}

func TestSpecDoesNotMatchCode(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, result := v.Validate()
	assert.Nil(t, err, "There is no error in execution: specs match code")
	assert.Equal(t, 1, len(result.SpecDefsNotInCode), "There is exactly 1 mismatch: specs do not match code")
}

func TestNoDirIgnoreFile(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, _ := v.Validate()
	assert.Error(t, err, "Error has occurred: "+err.Error())
}

func TestNoIgnoredAPIFile(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "", "./sample-app1", "./sample-app1/api-spec.yaml")
	err, _ := v.Validate()
	assert.Error(t, err, "Error has occurred: "+err.Error())
}

func TestNoSourceDir(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "", "./sample-app1/api-spec.yaml")
	err, _ := v.Validate()
	assert.Error(t, err, "Error has occurred: "+err.Error())
}

func TestNoAPIFile(t *testing.T) {
	v := validator.NewOpenAPISpecCodeDiffsValidator("./sample-app1/.dirignore", "./sample-app1/.filesignore", "./sample-app1/.linesignore", "./sample-app1/.apiignore", "./sample-app1", "")
	err, _ := v.Validate()
	assert.Error(t, err, "Error has occurred: "+err.Error())
}
