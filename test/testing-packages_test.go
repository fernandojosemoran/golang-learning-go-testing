package test

import (
	"testing"

	"github.com/fernandojosemoran/gotest/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFooerWithTestify(t *testing.T) {
	assert.Equal(t, "Foo", internal.Fooer(0), "0 is divisible by 3, should return Foo")
	assert.NotEqual(t, "Foo", internal.Fooer(1), "1 is not divisible by 3, should not return Foo")
}

func TestMapWithTestify(t *testing.T) {
	require.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "2", 3: "3"})
	assert.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "2"})
}
