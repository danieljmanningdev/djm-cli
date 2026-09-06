package djmcli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func Command(argv []string) error {
	if len(argv) < 2 {
		return errors.New("command required")
	}

	switch argv[1] {
	case "new":
		if len(argv) < 3 {
			return errors.New("project name required")
		}

		projectName := argv[2]

		clone := exec.Command(
			"git",
			"clone",
			"https://github.com/danieljmanningdev/go-starter-auth-app",
			projectName,
		)

		clone.Stdout = os.Stdout
		clone.Stderr = os.Stderr

		if err := clone.Run(); err != nil {
			return fmt.Errorf("clone starter: %w", err)
		}

		if err := os.RemoveAll(projectName + "/.git"); err != nil {
			return fmt.Errorf("remove starter git history: %w", err)
		}

		modEdit := exec.Command(
			"go",
			"mod",
			"edit",
			"-module=github.com/user/"+projectName,
		)

		modEdit.Dir = projectName

		if err := modEdit.Run(); err != nil {
			return fmt.Errorf("update module path: %w", err)
		}

		tidy := exec.Command("go", "mod", "tidy")
		tidy.Dir = projectName

		if err := tidy.Run(); err != nil {
			return fmt.Errorf("go mod tidy: %w", err)
		}

	case "add":
		// add feature

	case "dev":
		// run dev environment

	case "check":
		// run checks

	default:
		return fmt.Errorf("unknown command: %s", argv[1])
	}

	return nil
}
