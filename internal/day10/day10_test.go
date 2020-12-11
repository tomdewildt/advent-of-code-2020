package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidStream(t *testing.T) {
	stream := input.FromLiteral("abc\n20\n30\n40\n50\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 35, solution1, "Solution 1 should be 35")
	assert.Equal(t, 8, solution2, "Solution 2 should be 8")
	assert.Nil(t, err, "Error should be nil")
}
