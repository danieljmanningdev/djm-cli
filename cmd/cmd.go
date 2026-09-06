package cmd

import (
	"errors"
	"fmt"
)

func Command(argv []string) error {
	if len(argv) < 2 {
		return errors.New("command required")
	}

	switch argv[1] {
	case "new":
		return New(argv[2:])
	case "add":
		return Add(argv[2:])
	case "dev":
		return Dev(argv[2:])
	case "check":
		return Check(argv[2:])
	case "help", "--help", "-h":
		PrintHelp()
		return nil
	default:
		return fmt.Errorf("unknown command: %s", argv[1])
	}
}

func PrintHelp() {
	fmt.Println(`djm - Go-first web project tooling

Usage:
  djm new <name> [--module <path>]
  djm add <feature>
  djm dev
  djm check

Features:
  core
  auth
  security
  jsonld
  file-utils

Environment:
  DJM_MODULE_PREFIX   Default Go module prefix for new projects.
                      Defaults to github.com/danieljmanningdev.`)
}
