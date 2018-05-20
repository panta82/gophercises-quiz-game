package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

type Score struct {
	Total int
	Correct int
}

func executeQuiz(problems []Problem) Score {
	reader := bufio.NewReader(os.Stdin)
	score := Score{0, 0}
	i := 0

	for {
		if i >= len(problems) - 1 {
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
	}

	return score
}