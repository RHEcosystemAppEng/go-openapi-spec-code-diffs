package validator

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

// golangSourceFileReader Represents golang source file
type golangSourceFileReader struct {
	fileName     string
	ignoredLines *IgnoredLines
}

// NewGoLangSourceFileReader Returns a new golangSourceFileReader
func NewGoLangSourceFileReader(fileName string, ignoredLines *IgnoredLines) *golangSourceFileReader {
	return &golangSourceFileReader{fileName: fileName, ignoredLines: ignoredLines}
}

// GetAPILines Returns goSourceAPILine representing api paths found in golang source file
func (f *golangSourceFileReader) GetAPILines() (error, []*goSourceAPILine) {
	goFile, err := os.Open(f.fileName)
	if err != nil {
		log.Error().Msg("Error reading file: " + err.Error())
		return err, nil
	}
	fileScanner := bufio.NewScanner(goFile)
	fileScanner.Split(bufio.ScanLines)
	var lineNo int
	var apiLines []*goSourceAPILine
	for fileScanner.Scan() {
		lineNo++
		line := fileScanner.Text()
		if !slices.Contains(f.ignoredLines.ignoredLines, strings.TrimSpace(line)) {
			apiLine := NewGoSourceAPILine(f.fileName, line, lineNo, "", "")
			if apiLine.MatchesAPIEndpointRegEx() {
				apiLine.InferHttpMethod()
				apiLines = append(apiLines, apiLine)
			}
		}
	}

	return nil, apiLines
}
