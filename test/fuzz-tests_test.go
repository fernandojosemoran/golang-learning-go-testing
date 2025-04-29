package test

import (
	"testing"

	"github.com/fernandojosemoran/gotest/internal"
)

func FuzzFooer(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {
		internal.Fooer(a)
	})
}
