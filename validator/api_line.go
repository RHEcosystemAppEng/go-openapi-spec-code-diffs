package validator

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go/scanner"
	"go/token"
	"golang.org/x/exp/slices"
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
	uCaseLine := strings.ToUpper(g.fullLine)
	httpMethods := []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "HEAD", "PATCH"}
	for _, httpMethod := range httpMethods {
		uCaseHttpMethod := strings.ToUpper(httpMethod)
		if strings.Contains(uCaseLine, uCaseHttpMethod) {
			log.Debug().Msg(fmt.Sprintf("Potential API Definition with HTTP method '%s' found on line: %s", httpMethod, strings.TrimSpace(g.fullLine)))
			// Ensure that the http method is not defined in golang line with if, else, for, while etc. i.e. with a golang keyword
			if !g.lineHasGolangKeyword() {
				g.httpMethod = uCaseHttpMethod
				log.Debug().Msg(fmt.Sprintf("Added '%s' method for API Path '%s', which will be used to compare API paths defined in OpenAPI Specs File", g.httpMethod, g.apiPath))
			} else {
				log.Debug().Msg(fmt.Sprintf("May be this line is not an API definition as it has a golang keyword: %s", strings.TrimSpace(g.fullLine)))
			}
		}
	}
}

// lineHasGolangKeyword Determines if a goSourceLine contains a golang keyword
func (g *goSourceAPILine) lineHasGolangKeyword() bool {
	// source: https://go.dev/ref/spec#Keywords
	goKeyWords := []string{"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var"}

	src := []byte(g.fullLine)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil /* no error handler */, 0 /* Ignore comments */)

	// Repeated calls to Scan yield the token sequence found in the input.
	log.Debug().Msg(fmt.Sprintf("Potential api path further check of: %s", strings.TrimSpace(g.fullLine)))
	for {
		pos, tokType, lit := s.Scan()
		if tokType == token.EOF {
			break
		}
		log.Debug().Msg(fmt.Sprintf("%s %s %q", fset.Position(pos), tokType, lit))
		if slices.Contains(goKeyWords, tokType.String()) {
			return true
		}
	}
	return false
}

// MatchesAPIEndpointRegEx Checks whether a given golang source line is an api route definition based on a regular expression
func (g *goSourceAPILine) MatchesAPIEndpointRegEx() bool {
	API_DEF_REGEXP := `"\/.+"`
	var endpointRegExp = regexp.MustCompile(API_DEF_REGEXP)
	var isEndpointDefLine = endpointRegExp.MatchString(g.fullLine)
	if isEndpointDefLine {
		matchedString := endpointRegExp.FindAllString(g.fullLine, -1)[0]
		g.apiPath = strings.ReplaceAll(matchedString, "\"", "")
	}
	return isEndpointDefLine
}
