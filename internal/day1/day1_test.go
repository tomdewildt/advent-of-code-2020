package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidValue(t *testing.T) {
	stream := input.FromLiteral("abc\n979\n366\n299\n675\n1456")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolveNoPair(t *testing.T) {
	stream := input.FromLiteral("1720\n979\n366\n299\n675\n1456")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "pair", "Error should contain \"triple\"")
	assert.Contains(t, err.Error(), "not found", "Error should contain \"not found\"")

}

func TestSolveNoTriple(t *testing.T) {
	stream := input.FromLiteral("1721\n978\n366\n299\n675\n1456")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "triple", "Error should contain \"triple\"")
	assert.Contains(t, err.Error(), "not found", "Error should contain \"not found\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("1721\n979\n366\n299\n675\n1456")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 514579, solution1, "Solution 1 should be 514579")
	assert.Equal(t, 241861950, solution2, "Solution 2 should be 241861950")
	assert.Nil(t, err, "Error should be nil")
}
