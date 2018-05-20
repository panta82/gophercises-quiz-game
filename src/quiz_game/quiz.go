package main

import "fmt"

func main() {
	clArgs := parseCommandLineArgs()
	problems := loadProblemsFromFile(clArgs.Filename)

	for _, p := range problems {
		fmt.Println(fmt.Sprintf("%s -> %s", p.Question, p.Answer))
	}
}