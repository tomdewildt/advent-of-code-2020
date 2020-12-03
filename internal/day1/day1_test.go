package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidStream(t *testing.T) {
	stream := input.FromLiteral("abc\n20\n1000\n1010\n1020\n")

	solution1, solution2, err := solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolveNoPair(t *testing.T) {
	stream := input.FromLiteral("10\n20\n1000\n1010\n")

	solution1, solution2, err := solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "pair not found", "Error should contain \"pair not found\"")
}

func TestSolveNoTriple(t *testing.T) {
	stream := input.FromLiteral("20\n1000\n1010\n1020\n")

	solution1, solution2, err := solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "triple not found", "Error should contain \"triple not found\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("10\n20\n1000\n1010\n1020\n")

	solution1, solution2, err := solve(stream)

	assert.Equal(t, 1020000, solution1, "Solution 1 should be 1020000")
	assert.Equal(t, 10100000, solution2, "Solution 2 should be 10100000")
	assert.Nil(t, err, "Error should be nil")
}

func TestFindPairNotFound(t *testing.T) {
	value1, value2, err := findPair([]int{10, 20, 30, 40}, 10)

	assert.Equal(t, 0, value1, "Value 1 should be 0")
	assert.Equal(t, 0, value2, "Value 2 should be 0")
	assert.Equal(t, "pair not found", err.Error(), "Error should be \"pair not found\"")
}

func TestFindPair(t *testing.T) {
	value1, value2, err := findPair([]int{1, 2, 3, 4, 5}, 8)

	assert.Equal(t, 3, value1, "Value 1 should be 3")
	assert.Equal(t, 5, value2, "Value 2 should be 5")
	assert.Nil(t, err, "Error should be nil")
}

func TestFindTripleNotFound(t *testing.T) {
	value1, value2, value3, err := findTriple([]int{10, 20, 30, 40}, 10)

	assert.Equal(t, 0, value1, "Value 1 should be 0")
	assert.Equal(t, 0, value2, "Value 2 should be 0")
	assert.Equal(t, 0, value3, "Value 3 should be 0")
	assert.Equal(t, "triple not found", err.Error(), "Error should be \"triple not found\"")
}

func TestFindTriple(t *testing.T) {
	value1, value2, value3, err := findTriple([]int{1, 2, 3, 4, 5}, 8)

	assert.Equal(t, 3, value1, "Value 1 should be 3")
	assert.Equal(t, 4, value2, "Value 2 should be 4")
	assert.Equal(t, 1, value3, "Value 3 should be 1")
	assert.Nil(t, err, "Error should be nil")
}
