package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type runner struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func defaultRunner() runner {
	return runner{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

func (r runner) run(dir string, name string, args ...string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	command.Stdin = r.stdin
	command.Stdout = r.stdout
	command.Stderr = r.stderr

	if err := command.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}

	return nil
}

func commandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
