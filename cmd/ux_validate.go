package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func UXValidate(args []string) error {
	root := "docs/ux"

	var screens int
	var components int
	var flows int
	var tokens int

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
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
		case strings.HasSuffix(name, ".component.json"):
			components++
		case strings.HasSuffix(name, ".flow.json"):
			flows++
		case strings.HasSuffix(name, ".tokens.json"):
			tokens++
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
