package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"time"
)

type Score struct {
	Total int
	Correct int
	Timeout bool
}

func executeQuiz(problems []Problem, timeLimit int) Score {
	if timeLimit == 0 {
		// No timer, can just ask questions normally
		return doExecuteQuiz(problems, nil)
	}

	fmt.Printf("You have %d seconds to answer %d questions. Press ENTER when ready! > ", timeLimit, len(problems))
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n');
	if err != nil {
		fatal("Failed to read from stdin", err)
	}

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	scoreChan := make(chan Score)
	score := Score {}

	go doExecuteQuiz(problems, &scoreChan)

	loop:
	for {
		select {
		case <- timer.C:
			score.Timeout = true
			break loop
		case score = <- scoreChan:
			if score.Total == len(problems) {
				timer.Stop()
				break loop
			}
		}
	}

	return score
}

func doExecuteQuiz(problems []Problem, scoreChan *chan Score) Score {
	score := Score {}

	reader := bufio.NewReader(os.Stdin)
	i := 0

	for {
		if i >= len(problems) {
			break
		}

		fmt.Println("Q:  ", problems[i].Question)
		fmt.Print("A: > ")

		answer, err := reader.ReadString('\n')
		if err != nil {
			fatal("Input error", err)
		}

		answer = strings.Trim(answer, "\n\r \t")

		score.Total++
		if answer == problems[i].Answer {
			score.Correct++
		}

		i++

		if scoreChan != nil {
			*scoreChan <- score
		}
	}

	return score
}