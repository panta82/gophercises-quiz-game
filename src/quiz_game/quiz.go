package main

import "fmt"
import "encoding/csv"
import (
	"os"
	"io"
	"path/filepath"
)

const DefaultProblemsFilename = "problems.csv"

type Problem struct {
	Question string
	Answer string
}

func fatal(message string, err error) {
	fmt.Fprint(os.Stderr, "FATAL: " + message + "\n")
	if err != nil {
		fmt.Fprint(os.Stderr, "       " + err.Error() + "\n")
	}
	os.Exit(1)
}

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

func main() {
	problems := loadProblemsFromFile(getDefaultProblemsPath())

	for _, p := range problems {
		fmt.Println(fmt.Sprintf("%s -> %s", p.Question, p.Answer))
	}
}