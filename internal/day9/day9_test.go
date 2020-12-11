package day9

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
	stream, err := input.FromFile("../../assets/day9/day9.txt")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 88311122, solution1, "Solution 1 should be 88311122")
	assert.Equal(t, 13549369, solution2, "Solution 2 should be 13549369")
	assert.Nil(t, err, "Error should be nil")
}
