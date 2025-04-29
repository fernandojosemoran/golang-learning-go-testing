package test

import (
	"testing"

	"github.com/fernandojosemoran/gotest/internal"
)

func TestFooer(t *testing.T) {
	result := internal.Fooer(3)

	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want:%s", result, "Foo")
	}
}
