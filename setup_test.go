package main

import (
	"os"
	"testing"
)

func TestSomethingWhenGivenCondition(t *testing.T) {
	// Setup code for given condition goes here
	t.Run("Something has property one", func(t *testing.T) {
		// SubTest 1 code goes here
	})
	t.Run("Something has property two", func(t *testing.T) {
		// SubTest 2 code goes here
	})
	// Teardown code for given condition goes here
}

func TestMain(m *testing.M) {
	// Setup code goes here
	code := m.Run()
	// Teardown code goes here
	os.Exit(code)
}
