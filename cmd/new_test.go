package cmd

import "testing"

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
