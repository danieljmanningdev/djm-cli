package cmd

import (
	"reflect"
	"testing"
)

func TestAvailableFeatures(t *testing.T) {
	want := []string{
		"auth",
		"core",
		"file-utils",
		"jsonld",
		"security",
	}

	if got := availableFeatures(); !reflect.DeepEqual(got, want) {
		t.Fatalf("availableFeatures() = %v, want %v", got, want)
	}
}
