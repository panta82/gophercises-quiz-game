package main

import (
	"path/filepath"
	"os"
	"fmt"
	"encoding/csv"
	"io"
)

type Problem struct {
	Question string
	Answer string
}

const DefaultProblemsFilename = "problems.csv"

func getDefaultProblemsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fatal("Failed to determine directory of the executable", err)
	}

	return filepath.Join(dir, DefaultProblemsFilename)
}

func loadProblemsFromFile(path string) []Problem {
	file, err := os.Open(path)
	if err != nil {
		fatal(fmt.Sprintf("Failed to open %s", path), err)
	}

	var result []Problem
	var reader = csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal(fmt.Sprintf("Failed to parse %s", path), err)
		}

		result = append(result, Problem{row[0], row[1]})
	}

	return result
}