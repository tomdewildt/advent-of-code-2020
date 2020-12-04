package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolveInvalidStream(t *testing.T) {

}

func TestSolveInvalidMinCharacter(t *testing.T) {
	stream := input.FromLiteral("a-3 a: abcde\n1-3 b: cdefg\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolveInvalidMaxCharacter(t *testing.T) {
	stream := input.FromLiteral("1-a a: abcde\n1-3 b: cdefg\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("1-3 a: abcde\n1-3 b: cdefg\n1-3 a: abcde\n1-3 b: cdefg\n1-3 c: cccccc")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 2, solution1, "Solution 1 should be 2")
	assert.Equal(t, 2, solution2, "Solution 2 should be 2")
	assert.Nil(t, err, "Error should be nil")
}
