# OpenAPI Specs Validator

## Summary
A validation tool that compares OpenAPI specs vis-a-vis routes defined in code. Useful in scenarios where you want to ensure the OpenAPI specs and Code are in synch.

## What
This validator validates API Endpoints defined in golang source code with OpenAPI specifications and reports following differences.
* OpenAPI specs found in specifications but not found in golang source code
* Routes found in golang source code but not found in OpenAPI specs

## Why
Helps to keep OpenAPI specs and golang source code in synch

## How
* This tool takes golang source root directory which implements your API and OpenAPI specs file as inputs
* One can make use of .dirignore file to skip directories, which will be ignored to find API endpoints.
* You can also make use of .specignore file to ignore API paths to exclude from comparison such as /health/ready or /health/live

## Usage
### As shell command
* Format
```shell
openapi-spec-code-diffs 'path/to/openapi/specs/filename' 'path/to/golang/source/dir' 'path/to/ignored/directories/filename' 'path/to/ignored/paths/filename'
```
* Example
```shell
openapi-spec-code-diffs '~/example-service/openapi.yaml' '~/example-service' '~/example-service/.dirignore' '~/example-service/.specignore'
```

### As a package/library in golang source code
* Import the package/library
```go
import "github.com/RHEcosystemAppEng/openapi_spec_code_diffs/validator"
```

* Use the validator e.g. in test code as follows. Paths below are relative to the path of the test file from which tests are going to be executed from.
```go
func validateOpenAPISpecs(t *testing.T) {
	oasStaticValidator := validator.NewOpenAPISpecCodeDiffsValidator("./oasStaticValidator/.dirignore", "./oasStaticValidator/.specignore", "../../", "../../openapi.yaml")
	err, result := oasStaticValidator.Validate()

	assert.Nil(t, err, "No errors returned from openapi validation")
	assert.Equal(t, 0, len(result.SpecDefsNotInCode), "Found spec defs not implemented in code", len(result.SpecDefsNotInCode))
	assert.Equal(t, 0, len(result.CodeDefsNotSpec), "Found code defs not reflected in specs", len(result.CodeDefsNotSpec))
}
```
