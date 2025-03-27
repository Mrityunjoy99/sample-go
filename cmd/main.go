package main

import (
	"fmt"
	"os"

	"github.com/Mrityunjoy99/sample-go/cmd/cmdopts"
	"github.com/Mrityunjoy99/sample-go/migration"
	"github.com/Mrityunjoy99/sample-go/src/deployment/appserver"
)

func main() {
	args := cmdopts.Parse()

	const MaxArgLength = 2

	if len(os.Args) < MaxArgLength {
		fmt.Println("Invalid command")
		os.Exit(1)
	}

	command := os.Args[len(os.Args)-1]

	switch command {
	case "migrate":
		migration.Migrate(args)
	case "run":
		appserver.Start()
	default:
		panic("Invalid command")
	}
}
