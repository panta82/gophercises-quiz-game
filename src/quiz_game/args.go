package main

import (
	"flag"
	"fmt"
)

type CommandLineArgs struct {
	Filename string
	Limit int
	Shuffle bool
}

func parseCommandLineArgs() CommandLineArgs {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "QUIZ - a simple quiz game in go\n\n" +
			"Usage of quiz_game:");
		flag.PrintDefaults()
	}

	filename := flag.String("filename", "", "Filename of an alternative problems file (otherwise, uses the default)")
	limit := flag.Int("limit", 30, "Quiz time limit. Set 0 to disable")
	shuffle := flag.Bool("shuffle", false, "Shuffle problems before quiz")

	flag.Parse()

	result := CommandLineArgs{Filename: *filename, Limit: *limit, Shuffle: *shuffle}

	if result.Filename == "" {
		result.Filename = getDefaultProblemsPath()
	}

	return result
}