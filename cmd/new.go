package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	starterRepository   = "https://github.com/danieljmanningdev/go-starter-auth-app"
	starterModulePath   = "github.com/danieljmanningdev/go-starter-auth-app"
	defaultModulePrefix = "github.com/danieljmanningdev"
)

var projectNamePattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)

type newOptions struct {
	name       string
	modulePath string
}

func New(args []string) error {
	options, err := parseNewArgs(args)
	if err != nil {
		return err
	}

	return createProject(options, defaultRunner())
}

func parseNewArgs(args []string) (newOptions, error) {
	if len(args) == 0 {
		return newOptions{}, errors.New("project name required")
	}

	options := newOptions{name: args[0]}
	if err := validateProjectName(options.name); err != nil {
		return newOptions{}, err
	}

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--module":
			if i+1 >= len(args) {
				return newOptions{}, errors.New("--module requires a value")
			}

			options.modulePath = strings.TrimSpace(args[i+1])
			i++
		default:
			return newOptions{}, fmt.Errorf("unknown new option: %s", args[i])
		}
	}

	if options.modulePath == "" {
		prefix := strings.TrimSpace(os.Getenv("DJM_MODULE_PREFIX"))
		if prefix == "" {
			prefix = defaultModulePrefix
		}

		options.modulePath = strings.TrimRight(prefix, "/") + "/" + options.name
	}

	return options, nil
}

func validateProjectName(name string) error {
	if name == "" {
		return errors.New("project name required")
	}

	if name == "." || name == ".." || !projectNamePattern.MatchString(name) {
		return fmt.Errorf("invalid project name: %q", name)
	}

	return nil
}

func createProject(options newOptions, commands runner) (err error) {
	if _, statErr := os.Stat(options.name); statErr == nil {
		return fmt.Errorf("destination already exists: %s", options.name)
	} else if !errors.Is(statErr, os.ErrNotExist) {
		return fmt.Errorf("check destination: %w", statErr)
	}

	fmt.Printf("Creating %s from %s...\n", options.name, starterRepository)

	if err := commands.run("", "git", "clone", starterRepository, options.name); err != nil {
		return fmt.Errorf("clone starter: %w", err)
	}

	cleanup := true
	defer func() {
		if cleanup && err != nil {
			_ = os.RemoveAll(options.name)
		}
	}()

	if err := os.RemoveAll(filepath.Join(options.name, ".git")); err != nil {
		return fmt.Errorf("remove starter git history: %w", err)
	}

	if err := rewriteModuleImports(options.name, starterModulePath, options.modulePath); err != nil {
		return fmt.Errorf("rewrite starter imports: %w", err)
	}

	if err := commands.run(
		options.name,
		"go",
		"mod",
		"edit",
		"-module="+options.modulePath,
	); err != nil {
		return fmt.Errorf("update module path: %w", err)
	}

	if err := commands.run(options.name, "go", "mod", "tidy"); err != nil {
		return fmt.Errorf("tidy module: %w", err)
	}

	if err := commands.run(options.name, "git", "init", "-b", "main"); err != nil {
		return fmt.Errorf("initialise git repository: %w", err)
	}

	cleanup = false
	fmt.Printf("Created %s\nModule: %s\n", options.name, options.modulePath)
	return nil
}

func rewriteModuleImports(root, oldModule, newModule string) error {
	oldValue := []byte(oldModule)
	newValue := []byte(newModule)

	return filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if entry.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if !bytes.Contains(contents, oldValue) {
			return nil
		}

		info, err := entry.Info()
		if err != nil {
			return err
		}

		updated := bytes.ReplaceAll(contents, oldValue, newValue)
		if err := os.WriteFile(path, updated, info.Mode().Perm()); err != nil {
			return err
		}

		return nil
	})
}
