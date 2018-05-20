package main

import (
	"flag"
	"fmt"
)

type CommandLineArgs struct {
	Filename string
}

func parseCommandLineArgs() CommandLineArgs {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "QUIZ - a simple quiz game in go\n\n" +
			"Usage of quiz_game:");
		flag.PrintDefaults()
	}

	filename := flag.String("filename", "", "Filename of an alternative problems file (otherwise, uses the default)")

	flag.Parse()

	result := CommandLineArgs{Filename: *filename}

	if result.Filename == "" {
		result.Filename = getDefaultProblemsPath()
	}

	return result
}