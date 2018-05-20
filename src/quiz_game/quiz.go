package main

import "fmt"

func main() {
	clArgs := parseCommandLineArgs()
	problems := loadProblemsFromFile(clArgs.Filename)
	score := executeQuiz(problems)

	fmt.Println(fmt.Sprintf("You got %d out of %d correct", score.Correct, score.Total))
}