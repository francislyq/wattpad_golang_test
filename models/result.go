package models

import (
	"os"
	"strings"
)

type Result struct {
	FileName       string
	OffensiveScore int
}

// Get Results object
func (this Result) GetResults(files []os.FileInfo) ([]Result, error) {

	results := []Result{}

	for i := 0; i < len(files); i++ {
		this.FileName = files[i].Name()
		this.OffensiveScore = 0

		if strings.Contains(files[i].Name(), ".txt") {
			results = append(results, this)
		}
	}

	return results, nil
}
