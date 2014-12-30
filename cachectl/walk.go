package cachectl

import (
	"os"
	"path/filepath"
	"regexp"
)

func WalkPrintPagesStat(path string, re *regexp.Regexp) error {
	return filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if !info.Mode().IsRegular() {
				return nil
			}

			if re.MatchString(path) {
				PrintPagesStat(path, info.Size())
			}
			return nil
		})
}

func WalkPurgePages(path string, re *regexp.Regexp, rate float64) error {
	return filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if !info.Mode().IsRegular() {
				return nil
			}

			if re.MatchString(path) {
				RunPurgePages(path, info.Size(), rate)
			}
			return nil
		})
}
