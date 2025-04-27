package test

import (
	"testing"

	ubmo "github.com/fernandojosemoran/gotest/utilities"
)

func TestSum(test *testing.T) {
	result := ubmo.Sum(10, 15)
	expected := 25

	if result != expected {
		test.Errorf("Expected %d but got %d", expected, result)
	}

	result = ubmo.Sum(-10, 40)
	expected = 30

	if result != expected {
		test.Errorf("Expected %d but got %d", expected, result)
	}

	result = ubmo.Sum(0, -15)
	expected = -15

	if result != expected {
		test.Errorf("Expected %d but got %d", expected, result)
	}

	result = ubmo.Sum(-50, -15)
	expected = -65

	if result != expected {
		test.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestRest(test *testing.T) {
	result := ubmo.Rest(10, 15)
	expect := -5

	if result != expect {
		test.Errorf("Expect %d but got %d", expect, result)
	}

	result = ubmo.Rest(15, 10)
	expect = 5

	if result != expect {
		test.Errorf("Expect %d but got %d", expect, result)
	}

	result = ubmo.Rest(-50, -15)
	expect = -35

	if result != expect {
		test.Errorf("Expect %d but got %d", expect, result)
	}

	result = ubmo.Rest(-15, -50)
	expect = 35

	if result != expect {
		test.Errorf("Expect %d but got %d", expect, result)
	}
}
