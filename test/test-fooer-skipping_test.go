package test

import (
	"testing"

	"github.com/fernandojosemoran/gotest/internal"
)

func TestFooerSkipped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	result := internal.Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}
