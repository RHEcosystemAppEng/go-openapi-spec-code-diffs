package validator

import (
	"errors"
	"github.com/rs/zerolog/log"
	slices "golang.org/x/exp/slices"
	"io/fs"
	"path/filepath"
	"strings"
)

// GoSourceScanner Structure to specify root golang source directory and directories/files/lines to be ignored for scanning of API path definitions
type GoSourceScanner struct {
	dir                string
	ignoredDirectories *IgnoredDirectories
	ignoredFiles       *IgnoredFiles
	ignoredLines       *IgnoredLines
}

// NewGoSourceScanner Returns a new GoSourceScanner
func NewGoSourceScanner(dir string, ignoredDirectories *IgnoredDirectories, ignoredFiles *IgnoredFiles, ignoredLines *IgnoredLines) *GoSourceScanner {
	return &GoSourceScanner{dir: dir, ignoredDirectories: ignoredDirectories, ignoredFiles: ignoredFiles, ignoredLines: ignoredLines}
}

// ScanSourcesForAPIDefs Recursively walks the directory, analyzes golang source files to find API definitions ignoring directories found in .dirignore, ignoring files in .fileignore, ignoring lines in .lineignore
func (gss *GoSourceScanner) ScanSourcesForAPIDefs() (error, []*goSourceAPILine) {

	if strings.TrimSpace(gss.dir) == "" {
		return errors.New("No go source directory specified"), nil
	}

	var apiLines []*goSourceAPILine
	err := filepath.WalkDir(gss.dir,
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				if strings.HasSuffix(d.Name(), ".go") {
					log.Debug().Msg("Now scanning file: " + path)
					if slices.Contains(gss.ignoredFiles.ignoredFiles, path) {
						log.Debug().Msg("Ignoring scanning of go source file found in .ignoredfiles: " + path)
					} else {
						fileReader := NewGoLangSourceFileReader(path, gss.ignoredLines)
						err, fileAPILines := fileReader.GetAPILines()
						if err != nil {
							log.Error().Msg("Error reading API Defs in: " + d.Name() + " " + err.Error())
							return err
						}
						apiLines = append(apiLines, fileAPILines...)
					}
				}
			} else {
				if slices.Contains(gss.ignoredDirectories.ignoredDirectories, d.Name()) {
					return filepath.SkipDir
				}
			}
			return nil
		})
	if err != nil {
		log.Error().Msg("Impossible to traverse directories: " + gss.dir + " " + err.Error())
		return err, apiLines
	}

	return nil, apiLines
}
