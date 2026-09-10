package ux

import "testing"

func TestValidateScreen(t *testing.T) {
	validator, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	t.Run("valid screen", func(t *testing.T) {
		data := []byte(`{
			"id": "SCR-HOME",
			"name": "Home",
			"purpose": "Primary landing page"
		}`)

		if err := validator.ValidateScreen(data); err != nil {
			t.Fatalf("ValidateScreen() error = %v", err)
		}
	})

	t.Run("invalid screen id", func(t *testing.T) {
		data := []byte(`{
			"id": "BAD",
			"name": "Home",
			"purpose": "Primary landing page"
		}`)

		if err := validator.ValidateScreen(data); err == nil {
			t.Fatal("ValidateScreen() expected error, got nil")
		}
	})
}

func TestValidateComponent(t *testing.T) {
	validator, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	t.Run("valid component", func(t *testing.T) {
		data := []byte(`{
			"id": "CMP-HEADER",
			"name": "Header",
			"purpose": "Primary site navigation",
			"type": "navigation"
		}`)

		if err := validator.ValidateComponent(data); err != nil {
			t.Fatalf("ValidateComponent() error = %v", err)
		}
	})

	t.Run("missing component type", func(t *testing.T) {
		data := []byte(`{
			"id": "CMP-HEADER",
			"name": "Header",
			"purpose": "Primary site navigation"
		}`)

		if err := validator.ValidateComponent(data); err == nil {
			t.Fatal("ValidateComponent() expected error, got nil")
		}
	})
}

func TestValidateFlow(t *testing.T) {
	validator, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	t.Run("valid flow", func(t *testing.T) {
		data := []byte(`{
			"id": "FLOW-AUTH",
			"name": "Authentication",
			"purpose": "Authenticate a user",
			"steps": [
				{
					"id": "LOGIN",
					"type": "screen",
					"name": "Login"
				}
			]
		}`)

		if err := validator.ValidateFlow(data); err != nil {
			t.Fatalf("ValidateFlow() error = %v", err)
		}
	})

	t.Run("flow requires steps", func(t *testing.T) {
		data := []byte(`{
			"id": "FLOW-AUTH",
			"name": "Authentication",
			"purpose": "Authenticate a user",
			"steps": []
		}`)

		if err := validator.ValidateFlow(data); err == nil {
			t.Fatal("ValidateFlow() expected error, got nil")
		}
	})
}

func TestValidateTokens(t *testing.T) {
	validator, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	t.Run("valid tokens", func(t *testing.T) {
		data := []byte(`{
			"color": {
				"primary": {
					"value": "#000000",
					"type": "color"
				}
			}
		}`)

		if err := validator.ValidateTokens(data); err != nil {
			t.Fatalf("ValidateTokens() error = %v", err)
		}
	})

	t.Run("token requires type", func(t *testing.T) {
		data := []byte(`{
			"color": {
				"primary": {
					"value": "#000000"
				}
			}
		}`)

		if err := validator.ValidateTokens(data); err == nil {
			t.Fatal("ValidateTokens() expected error, got nil")
		}
	})
}
