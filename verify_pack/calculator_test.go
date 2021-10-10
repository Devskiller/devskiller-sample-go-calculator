package verify_pack

import (
	"math"
	"testing"

	"github.com/devskiller/task/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAddition_Assessment(t *testing.T) {
	total := calculator.Add(-1, -1)
	assert.EqualValues(t, -2, total, "add result was incorrect, got: %v, expected: %v.", total, -2)
}

func TestSubtraction_Assessment(t *testing.T) {
	total := calculator.Subtract(2, -1)
	assert.EqualValues(t, 3, total, "subtraction result was incorrect, got: %v, expected: %v.", total, 3)
}

func TestMultiplicationByZero_Assessment(t *testing.T) {
	total := calculator.Multiply(2, 0)
	assert.EqualValues(t, 0, total, "multiplication by 0 result was incorrect, got: %v, expected: %v.", total, 0)
}

func TestDivisionByZero_Assessment(t *testing.T) {
	total := calculator.Divide(4, 0.0)
	assert.EqualValues(t, math.Inf(1), total, "division by 0 result was incorrect, got: %v, expected: %v.", total, 2)
}
