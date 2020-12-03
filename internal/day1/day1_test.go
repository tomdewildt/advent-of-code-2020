package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidStream(t *testing.T) {
	stream := input.FromLiteral("abc\n20\n1000\n1010\n1020\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolveNoPair(t *testing.T) {
	stream := input.FromLiteral("10\n20\n1000\n1010\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "pair not found", "Error should contain \"pair not found\"")
}

func TestSolveNoTriple(t *testing.T) {
	stream := input.FromLiteral("20\n1000\n1010\n1020\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "triple not found", "Error should contain \"triple not found\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("10\n20\n1000\n1010\n1020\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 1020000, solution1, "Solution 1 should be 1020000")
	assert.Equal(t, 10100000, solution2, "Solution 2 should be 10100000")
	assert.Nil(t, err, "Error should be nil")
}
