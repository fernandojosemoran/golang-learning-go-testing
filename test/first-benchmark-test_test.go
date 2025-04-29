package test

import (
	"testing"

	"github.com/fernandojosemoran/gotest/internal"
)

func BenchMarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		internal.Fooer(i)
	}
}
