# OpenAPI Specs Validator

## Summary
A validation tool that compares OpenAPI specs vis-a-vis routes (e.g. /api/v1/customer/:id) defined in golang source code. Useful in scenarios where you want to ensure the OpenAPI specs and Code are in synch.

## What
This validator validates API Endpoints defined in golang source code with OpenAPI specifications and reports following differences.
* OpenAPI specs found in specifications but not found in golang source code
* Routes found in golang source code but not found in OpenAPI specs

## Why
* Helps to keep OpenAPI specs and golang source code in synch.
* Do not forget to implement paths defined in OpenAPI specs in golang source code
* Do not forget to include APIs implemented in golang source code in OpenAPI specs in

## How
* This tool takes golang source root directory which implements your API and OpenAPI specs file as inputs
* The tool essentially builds two lists: 
  1. List of routes/paths found in golang source code 
  2. List of routes/paths defined in your OpenAPI specs file
  * Both the lists should match if not the tool will report differences
* The tool uses regular expressions to find paths defined in golang source code e.g. "/health/ready", "/users" etc.
  * Once such a path is found, the corresponding line is scanned to look for a httpmethod such as GET, PUT, POST, DELETE, HEAD, OPTIONS and PATCH.
  * If the httpmethod is found on the line then tool considers the path found in the line as a valid route/path definition.
  * Thus following code is considered a valid route/path definition by the tool
```go
route := routeRegistration{"PUT", "/user/:id/admin/:isAdmin", handlers.SetAdminStatus}
```
  * However the following line is not considered a valid route/path definition as this is the line having a golang keyword 'if'
```go
if "/this/is/not/considered/a/path/definition" == "DELETE" {
```
  * You can make use of the following ignore elements while scanning golang source code, to further fine tune the run of the tool. 

| Element       | Description |
|------------------|-------------|
| ignoredDirsFile  | Directories to be ignored when scanning golang source code |
| ignoredFilesFile | Files to be ignored |
| ignoredLinesFile | Lines to be ignored |
| ignoredPathsFile | API Paths to be ignored from comparison such as /health/ready or /health/live |

## Why the ignored elements above are needed?
* The tool recursively works on .go files found in golang source code directory.
* You can slightly improve the performance by specifying directories which typically do not contain .go files such as bin, out and .git e.g.
* You can use ignoredFilesFile to further exclude specific files from scanning.
* If the ignored directories and files are not sufficient, if tool confuses certain code lines to contain valid paths/routes but if they are not valid paths/routes, simply copy those line in ignoredLinesFile and those lines will be ignored from scan.

## Usage
### As shell command
* The command line version makes use of named arguments. 
* At any point to see the help for the command use the following
```shell
openapi_spec_code_diffs --help
```
* Format
```shell
openapi_spec_code_diffs /
    -openAPISpecsFile 'path/to/openapi/specs/filename' /
    -goSourceDir 'path/to/golang/source/dir' /
    -ignoredDirsFile 'path/to/ignored/directories/filename' / 
    -ignoredFilesFile 'path/to/ignored/files/filename' /
    -ignoredLinesFile 'path/to/ignored/lines/filename' / 
    -ignoredPathsFile 'path/to/ignored/paths/filename' / 
    -logLevel 'log-level can be one of: disabled, info, debug, error'
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

## Running tests
* Use the standard go command to run tests as follows, which are divided as positive and negative tests:
```shell
 go test -v ./tests/positive ./tests/negative
```