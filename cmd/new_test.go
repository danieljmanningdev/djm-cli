package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "good-app"},
		{name: "good_app"},
		{name: "good.app"},
		{name: "app123"},
		{name: "", wantErr: true},
		{name: ".", wantErr: true},
		{name: "..", wantErr: true},
		{name: "bad app", wantErr: true},
		{name: "/tmp/app", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProjectName(tt.name)
			if tt.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestParseNewArgsDefaultModule(t *testing.T) {
	t.Setenv("DJM_MODULE_PREFIX", "github.com/example")

	options, err := parseNewArgs([]string{"good-app"})
	if err != nil {
		t.Fatalf("parseNewArgs returned error: %v", err)
	}

	if options.name != "good-app" {
		t.Fatalf("name = %q, want %q", options.name, "good-app")
	}

	if options.modulePath != "github.com/example/good-app" {
		t.Fatalf("modulePath = %q, want %q", options.modulePath, "github.com/example/good-app")
	}
}

func TestParseNewArgsExplicitModule(t *testing.T) {
	options, err := parseNewArgs([]string{
		"good-app",
		"--module",
		"example.com/team/good-app",
	})
	if err != nil {
		t.Fatalf("parseNewArgs returned error: %v", err)
	}

	if options.modulePath != "example.com/team/good-app" {
		t.Fatalf("modulePath = %q, want %q", options.modulePath, "example.com/team/good-app")
	}
}

func TestRewriteModuleImports(t *testing.T) {
	root := t.TempDir()
	goFile := filepath.Join(root, "main.go")
	textFile := filepath.Join(root, "README.md")

	originalGo := `package main

import "github.com/danieljmanningdev/go-starter-auth-app/internal/auth"
`
	originalText := "github.com/danieljmanningdev/go-starter-auth-app"

	if err := os.WriteFile(goFile, []byte(originalGo), 0o644); err != nil {
		t.Fatalf("write go file: %v", err)
	}
	if err := os.WriteFile(textFile, []byte(originalText), 0o644); err != nil {
		t.Fatalf("write text file: %v", err)
	}

	if err := rewriteModuleImports(
		root,
		starterModulePath,
		"github.com/example/my-app",
	); err != nil {
		t.Fatalf("rewriteModuleImports returned error: %v", err)
	}

	updatedGo, err := os.ReadFile(goFile)
	if err != nil {
		t.Fatalf("read go file: %v", err)
	}

	if strings.Contains(string(updatedGo), starterModulePath) {
		t.Fatalf("starter module path still present in Go source: %s", updatedGo)
	}
	if !strings.Contains(string(updatedGo), "github.com/example/my-app/internal/auth") {
		t.Fatalf("new module path missing from Go source: %s", updatedGo)
	}

	updatedText, err := os.ReadFile(textFile)
	if err != nil {
		t.Fatalf("read text file: %v", err)
	}
	if string(updatedText) != originalText {
		t.Fatalf("non-Go file changed: %q", updatedText)
	}
}
