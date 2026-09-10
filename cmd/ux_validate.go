package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/danieljmanningdev/djm-cli/internal/ux"
)

func UXValidate(args []string) error {
	root := "docs/ux"

	validator, err := ux.NewValidator()
	if err != nil {
		return fmt.Errorf("load UX schemas: %w", err)
	}

	var screens int
	var components int
	var flows int
	var tokens int

	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		name := entry.Name()

		switch {
		case strings.HasSuffix(name, ".screen.json"):
			screens++

			if err := validateScreen(path, validator); err != nil {
				return err
			}

		case strings.HasSuffix(name, ".component.json"):
			components++

			if err := validateComponent(path, validator); err != nil {
				return err
			}

		case strings.HasSuffix(name, ".flow.json"):
			flows++

			if err := validateFlow(path, validator); err != nil {
				return err
			}

		case strings.HasSuffix(name, ".tokens.json"):
			tokens++

			if err := validateTokens(path, validator); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("scan UX documents: %w", err)
	}

	total := screens + components + flows + tokens

	fmt.Printf("Found %d UX documents\n", total)
	fmt.Printf("%d screens\n", screens)
	fmt.Printf("%d components\n", components)
	fmt.Printf("%d flows\n", flows)
	fmt.Printf("%d token files\n", tokens)

	return nil
}

func validateScreen(path string, validator *ux.Validator) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	if err := validator.ValidateScreen(data); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	fmt.Printf("✓ %s\n", path)

	return nil
}

func validateComponent(path string, validator *ux.Validator) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	if err := validator.ValidateComponent(data); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	fmt.Printf("✓ %s\n", path)

	return nil
}

func validateFlow(path string, validator *ux.Validator) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	if err := validator.ValidateFlow(data); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	fmt.Printf("✓ %s\n", path)

	return nil
}

func validateTokens(path string, validator *ux.Validator) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	if err := validator.ValidateTokens(data); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	fmt.Printf("✓ %s\n", path)

	return nil
}
