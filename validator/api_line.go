package validator

import (
	"regexp"
	"strings"
)

// goSourceAPILine Struct that represents an API route definition in golang source code.
type goSourceAPILine struct {
	fileName   string
	fullLine   string
	lineNum    int
	httpMethod string
	apiPath    string
}

// NewGoSourceAPILine Returns a new goSourceAPILine
func NewGoSourceAPILine(fileName string, fullLine string, lineNum int, httpMethod string, apiPath string) *goSourceAPILine {
	return &goSourceAPILine{
		fileName:   fileName,
		fullLine:   fullLine,
		lineNum:    lineNum,
		httpMethod: httpMethod,
		apiPath:    apiPath,
	}
}

// InferHttpMethod Analyzes a give line and infers http method such as GET, PUT, POST, DELETE, OPTIONS, HEAD OR PATCH
func (g *goSourceAPILine) InferHttpMethod() {
	upperCaseLine := strings.ToUpper(g.fullLine)
	if strings.Contains(upperCaseLine, "\"GET\"") || strings.Contains(upperCaseLine, ".GET") {
		g.httpMethod = "GET"
	} else if strings.Contains(upperCaseLine, "\"PUT\"") || strings.Contains(upperCaseLine, ".PUT") {
		g.httpMethod = "PUT"
	} else if strings.Contains(upperCaseLine, "\"POST\"") || strings.Contains(upperCaseLine, ".POST") {
		g.httpMethod = "POST"
	} else if strings.Contains(upperCaseLine, "\"DELETE\"") || strings.Contains(upperCaseLine, ".DELETE") {
		g.httpMethod = "DELETE"
	} else if strings.Contains(upperCaseLine, "\"OPTIONS\"") || strings.Contains(upperCaseLine, ".OPTIONS") {
		g.httpMethod = "OPTIONS"
	} else if strings.Contains(upperCaseLine, "\"PATCH\"") || strings.Contains(upperCaseLine, ".PATCH") {
		g.httpMethod = "PATCH"
	} else if strings.Contains(upperCaseLine, "\"HEAD\"") || strings.Contains(upperCaseLine, ".HEAD") {
		g.httpMethod = "HEAD"
	}
}

// IsAPIEndpointDefLine Checks whether a given golang source line is an api route definition based on regular expression
func (g *goSourceAPILine) IsAPIEndpointDefLine() bool {
	API_DEF_REGEXP := `"\/.+"`
	var endpointRegExp = regexp.MustCompile(API_DEF_REGEXP)
	var isEndpointDefLine = endpointRegExp.MatchString(g.fullLine)
	if isEndpointDefLine {
		matchedString := endpointRegExp.FindAllString(g.fullLine, -1)[0]
		g.apiPath = strings.ReplaceAll(matchedString, "\"", "")
	}
	return isEndpointDefLine
}
