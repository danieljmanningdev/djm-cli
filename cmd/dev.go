package cmd

import "errors"

func Dev(args []string) error {
	if len(args) != 0 {
		return errors.New("dev does not accept arguments")
	}

	return defaultRunner().run("", "go", "run", "./cmd/server")
}
