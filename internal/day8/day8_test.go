package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidInstruction(t *testing.T) {
	stream := input.FromLiteral("nop +0\nacc +1\njmp +2\nnop +0\nnop +0\njmp -4\nabc +1")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid operation", "Error should contain \"invalid operation\"")
}

func TestSolveInvalidArgument(t *testing.T) {
	stream := input.FromLiteral("nop +0\nacc +1\njmp +2\nnop +0\nnop +0\njmp -4\nacc abc")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("nop +0\nacc +1\njmp +2\nnop +0\nnop +0\njmp -4\nacc +1")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 1, solution1, "Solution 1 should be 1")
	assert.Equal(t, 2, solution2, "Solution 2 should be 2")
	assert.Nil(t, err, "Error should be nil")
}
