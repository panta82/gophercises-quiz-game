package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	clArgs := parseCommandLineArgs()
	problems := loadProblemsFromFile(clArgs.Filename, clArgs.Shuffle)
	score := executeQuiz(problems, clArgs.Limit)

	scoreMsg := fmt.Sprintf("You got %d out of %d correct", score.Correct, len(problems))
	if score.Timeout {
		scoreMsg += fmt.Sprintf(" (ran out of time after %d questions)", score.Total)
	}
	fmt.Println(scoreMsg)
}