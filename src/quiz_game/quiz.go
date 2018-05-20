package main

import "fmt"

func main() {
	clArgs := parseCommandLineArgs()
	problems := loadProblemsFromFile(clArgs.Filename)
	score := executeQuiz(problems, clArgs.Limit)

	scoreMsg := fmt.Sprintf("You got %d out of %d correct", score.Correct, len(problems))
	if score.Timeout {
		scoreMsg += fmt.Sprintf(" (ran out of time after %d questions)", score.Total)
	}
	fmt.Println(scoreMsg)
}