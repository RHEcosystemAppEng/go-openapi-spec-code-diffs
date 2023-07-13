package validator

import (
	"github.com/rs/zerolog/log"
	slices "golang.org/x/exp/slices"
	"io/fs"
	"path/filepath"
	"strings"
)

// GoSourceDirectory Structure to specify root golang source directory and directories to be ignored
type GoSourceDirectory struct {
	dir                string
	ignoredDirectories *IgnoredDirectories
}

// NewGoSourceDirectory Returns a new GoSourceDirectory
func NewGoSourceDirectory(dir string, ignoredDirectories *IgnoredDirectories) *GoSourceDirectory {
	return &GoSourceDirectory{dir: dir, ignoredDirectories: ignoredDirectories}
}

// WalkDirectory Recursively walks the directory, analyzes golang source files to find API definitions ignoring directories found in .dirignore
func (gsd *GoSourceDirectory) WalkDirectory() (error, []*goSourceAPILine) {
	var apiLines []*goSourceAPILine
	err := filepath.WalkDir(gsd.dir,
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				if strings.HasSuffix(d.Name(), ".go") {
					fileReader := NewGoLangSourceFileReader(path)
					err, fileAPILines := fileReader.GetAPILines()
					if err != nil {
						log.Error().Msg("Error reading API Defs in: " + d.Name() + " " + err.Error())
						return err
					}
					apiLines = append(apiLines, fileAPILines...)
				}
			} else {
				if slices.Contains(gsd.ignoredDirectories.ignoredDirectories, d.Name()) {
					return filepath.SkipDir
				}
			}
			return nil
		})
	if err != nil {
		log.Error().Msg("Impossible to traverse directories: " + gsd.dir + " " + err.Error())
		return err, apiLines
	}

	return nil, apiLines
}
