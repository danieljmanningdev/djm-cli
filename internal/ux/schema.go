package ux

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

//go:embed schemas/*.json definitions/*.json
var SchemaFS embed.FS

type Validator struct {
	screen    *jsonschema.Schema
	component *jsonschema.Schema
	flow      *jsonschema.Schema
	tokens    *jsonschema.Schema
}

func NewValidator() (*Validator, error) {
	compiler := jsonschema.NewCompiler()

	definitions, err := fs.Glob(SchemaFS, "definitions/*.json")
	if err != nil {
		return nil, fmt.Errorf("find schema definitions: %w", err)
	}

	for _, path := range definitions {
		if err := addSchemaResource(compiler, path); err != nil {
			return nil, err
		}
	}

	schemaPaths := []string{
		"schemas/screen.schema.json",
		"schemas/component.schema.json",
		"schemas/flow.schema.json",
		"schemas/tokens.schema.json",
	}

	for _, path := range schemaPaths {
		if err := addSchemaResource(compiler, path); err != nil {
			return nil, err
		}
	}

	screen, err := compiler.Compile("schemas/screen.schema.json")
	if err != nil {
		return nil, fmt.Errorf("compile screen schema: %w", err)
	}

	component, err := compiler.Compile("schemas/component.schema.json")
	if err != nil {
		return nil, fmt.Errorf("compile component schema: %w", err)
	}

	flow, err := compiler.Compile("schemas/flow.schema.json")
	if err != nil {
		return nil, fmt.Errorf("compile flow schema: %w", err)
	}

	tokens, err := compiler.Compile("schemas/tokens.schema.json")
	if err != nil {
		return nil, fmt.Errorf("compile tokens schema: %w", err)
	}

	return &Validator{
		screen:    screen,
		component: component,
		flow:      flow,
		tokens:    tokens,
	}, nil
}

func addSchemaResource(compiler *jsonschema.Compiler, path string) error {
	data, err := SchemaFS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read embedded schema %s: %w", path, err)
	}

	document, err := jsonschema.UnmarshalJSON(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("parse embedded schema %s: %w", path, err)
	}

	if err := compiler.AddResource(path, document); err != nil {
		return fmt.Errorf("register embedded schema %s: %w", path, err)
	}

	return nil
}

func validate(schema *jsonschema.Schema, data []byte) error {
	document, err := jsonschema.UnmarshalJSON(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	if err := schema.Validate(document); err != nil {
		return err
	}

	return nil
}

func (v *Validator) ValidateScreen(data []byte) error {
	return validate(v.screen, data)
}

func (v *Validator) ValidateComponent(data []byte) error {
	return validate(v.component, data)
}

func (v *Validator) ValidateFlow(data []byte) error {
	return validate(v.flow, data)
}

func (v *Validator) ValidateTokens(data []byte) error {
	return validate(v.tokens, data)
}
