package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

const instructions = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestSolveInvalidInstruction(t *testing.T) {
	stream := input.FromLiteral("abc +1")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid operation", "Error should contain \"invalid operation\"")
}

func TestSolveInvalidArgument(t *testing.T) {
	stream := input.FromLiteral("acc abc")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral(instructions)

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 5, solution1, "Solution 1 should be 5")
	assert.Equal(t, 8, solution2, "Solution 2 should be 8")
	assert.Nil(t, err, "Error should be nil")
}
