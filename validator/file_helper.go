package validator

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"os"
)

// ReadFile Reads file contents and returns its contents as an array of strings.
func ReadFile(filename string) (error, []string) {
	var contents []string
	goFile, err := os.Open(filename)
	if err != nil {
		log.Error().Msg("Error reading contents of file: " + filename + " " + err.Error())
		return err, nil
	}
	fileScanner := bufio.NewScanner(goFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		contents = append(contents, line)
	}

	return nil, contents
}
