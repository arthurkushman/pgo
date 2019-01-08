package pgo

import (
	"fmt"
	"os"
)

func isOk(ok bool, msg string, args ...interface{}) {
	if !ok {
		printError(msg, args)
	}
}

func printError(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	if os.Getenv("PGO_ENV") != "dev" {
		os.Exit(1)
	}
}