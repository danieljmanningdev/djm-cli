package cmd

import (
	"errors"
	"fmt"
)

func UX(args []string) error {
	if len(args) == 0 {
		return errors.New("ux command required")
	}

	switch args[0] {
	case "validate":
		return UXValidate(args[1:])
	default:
		return fmt.Errorf("unknown ux command: %s", args[0])
	}
}
