package cmd

import (
	"errors"
	"fmt"
	"runtime/debug"
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
	case "version", "--version", "-v":
		PrintVersion()
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
  djm version

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

func PrintVersion() {
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" || info.Main.Version == "(devel)" {
		fmt.Println("djm dev")
		return
	}

	fmt.Printf("djm %s\n", info.Main.Version)
}
