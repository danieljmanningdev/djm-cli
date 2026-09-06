package cmd

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var featureModules = map[string]string{
	"auth":       "github.com/danieljmanningdev/go-web-auth@latest",
	"core":       "github.com/danieljmanningdev/go-web-core@latest",
	"file-utils": "github.com/danieljmanningdev/go-file-utils@latest",
	"jsonld":     "github.com/danieljmanningdev/go-jsonld-schema@latest",
	"security":   "github.com/danieljmanningdev/go-web-security@latest",
}

func Add(args []string) error {
	if len(args) == 0 {
		return errors.New("feature required")
	}

	if len(args) > 1 {
		return errors.New("add accepts one feature at a time")
	}

	feature := strings.ToLower(strings.TrimSpace(args[0]))
	module, ok := featureModules[feature]
	if !ok {
		return fmt.Errorf("unknown feature %q; available: %s", feature, strings.Join(availableFeatures(), ", "))
	}

	fmt.Printf("Adding %s...\n", feature)
	if err := defaultRunner().run("", "go", "get", module); err != nil {
		return fmt.Errorf("add %s: %w", feature, err)
	}

	fmt.Printf("Added %s\n", feature)
	return nil
}

func availableFeatures() []string {
	features := make([]string, 0, len(featureModules))
	for feature := range featureModules {
		features = append(features, feature)
	}

	sort.Strings(features)
	return features
}
