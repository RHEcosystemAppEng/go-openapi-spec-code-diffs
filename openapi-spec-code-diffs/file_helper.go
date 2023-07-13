package openapi_spec_code_diffs

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"os"
)

// ReadFile Reads file contents and returns its contents as an array of strings.
func ReadFile(filename string) (error, []string) {
	var contents []string
	f := golangSourceFileReader{fileName: filename}
	goFile, err := os.Open(f.fileName)
	if err != nil {
		log.Error().Msg("Error reading contents of file: " + filename + " " + err.Error())
		return err, nil
	}
	fileScanner := bufio.NewScanner(goFile)
	fileScanner.Split(bufio.ScanLines)
	var lineNo int
	for fileScanner.Scan() {
		lineNo++
		line := fileScanner.Text()
		contents = append(contents, line)
	}

	return nil, contents
}
