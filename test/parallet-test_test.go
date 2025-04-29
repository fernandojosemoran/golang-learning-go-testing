package test

import (
	"testing"

	"github.com/fernandojosemoran/gotest/internal"
)

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in parallel", func(t *testing.T) {
		t.Parallel()
		result := internal.Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s", result, "Foo")
		}
	})

	t.Run("Test 7 in parallel", func(t *testing.T) {
		t.Parallel()
		result := internal.Fooer(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s", result, "7")
		}
	})
}
