package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	total := Add(2, 3)
	assert.EqualValues(t, 5, total, "add result was incorrect, got: %v, expected: %v.", total, 5)
}

func TestSubtraction(t *testing.T) {
	total := Subtract(2, 3)
	assert.EqualValues(t, -1, total, "subtraction result was incorrect, got: %v, expected: %v.", total, -1)
}

func TestMultiplication(t *testing.T) {
	total := Multiply(2, 3)
	assert.EqualValues(t, 6, total, "multiplication result was incorrect, got: %v, expected: %v.", total, 6)
}

func TestDivision(t *testing.T) {
	total := Divide(4, 2)
	assert.EqualValues(t, 2, total, "division result was incorrect, got: %v, expected: %v.", total, 2)
}
