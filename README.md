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
* Format
```shell
openapi-spec-code-diffs 'path/to/openapi/specs/filename' 'path/to/golang/source/dir' 'path/to/ignored/directories/filename' 'path/to/ignored/paths/filename'
```
* Example
```shell
openapi-spec-code-diffs '~/example-service/openapi.yaml' '~/example-service' '~/example-service/.dirignore' '~/example-service/.specignore'
```