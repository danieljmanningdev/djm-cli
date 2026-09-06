package cmd

import (
	"errors"
	"fmt"
)

type checkStep struct {
	name string
	args []string
}

func Check(args []string) error {
	if len(args) != 0 {
		return errors.New("check does not accept arguments")
	}

	steps := []checkStep{
		{name: "gofmt", args: []string{"-w", "."}},
		{name: "go", args: []string{"vet", "./..."}},
		{name: "go", args: []string{"test", "./..."}},
	}

	commands := defaultRunner()
	for _, step := range steps {
		fmt.Printf("Running %s %v...\n", step.name, step.args)
		if err := commands.run("", step.name, step.args...); err != nil {
			return err
		}
	}

	if commandAvailable("govulncheck") {
		fmt.Println("Running govulncheck ./...")
		if err := commands.run("", "govulncheck", "./..."); err != nil {
			return err
		}
	} else {
		fmt.Println("Skipping govulncheck: command not installed")
	}

	if commandAvailable("git") {
		fmt.Println("Running git diff --check")
		if err := commands.run("", "git", "diff", "--check"); err != nil {
			return err
		}
	}

	fmt.Println("Checks passed")
	return nil
}
