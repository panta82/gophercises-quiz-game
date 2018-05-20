package main

import (
	"fmt"
	"os"
)

func fatal(message string, err error) {
	fmt.Fprint(os.Stderr, "FATAL: " + message + "\n")
	if err != nil {
		fmt.Fprint(os.Stderr, "       " + err.Error() + "\n")
	}
	os.Exit(1)
}
