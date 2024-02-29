package cmdopts

import (
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
	MigrateAll   bool
	MigrateTo    string
	RollbackLast bool
	RollbackTo   string
	CommsWorker  bool
}

func getHelpText() string {
	return `Usage:
	main <command> [arguments]

The commands are:
  start	starts the server
  migrate performs DB migration.

The arguments are:
`
}

func Parse() *Arguments {
	a := Arguments{}
	flag.BoolVar(
		&a.MigrateAll,
		MigrateAllOpts.ToString(),
		false,
		"migrates to the latest version",
	)

	flag.StringVar(
		&a.MigrateTo,
		MigrateToOpts.ToString(),
		"",
		"migrate to the given version",
	)

	flag.BoolVar(
		&a.RollbackLast,
		MigrateLastOpts.ToString(),
		false,
		"rollback the last version",
	)

	flag.StringVar(
		&a.RollbackTo,
		RollbackToOpts.ToString(),
		"",
		"rollback to the given vesion",
	)

	flag.BoolVar(
		&a.CommsWorker,
		CommsWorker.ToString(),
		false,
		"rollback the last version",
	)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, getHelpText())
		flag.PrintDefaults()
	}

	flag.Parse()

	return &a
}

func IsFlagPassed(name string) bool {
	found := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
