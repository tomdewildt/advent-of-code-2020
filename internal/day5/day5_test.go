package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidSeat(t *testing.T) {
	stream := input.FromLiteral("ABCDEFGHIJ\nFFFBBFFRRL\nFFFBBFBLLL\nFFFBBFBLLR\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("FFFBBFFRLR\nFFFBBFFRRL\nFFFBBFBLLL\nFFFBBFBLLR\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 105, solution1, "Solution 1 should be 105")
	assert.Equal(t, 103, solution2, "Solution 2 should be 103")
	assert.Nil(t, err, "Error should be nil")
}
