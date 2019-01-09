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
	fmt.Printf(msg, []interface{}(args)) // strange fix but it didn't work on go version go1.11.4 darwin/amd64 with args...
	if os.Getenv("PGO_ENV") != "dev" {
		os.Exit(1)
	}
}
