package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("abc\n\na\nb\nc\n\nab\nac\n\na\na\na\n\na")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 11, solution1, "Solution 1 should be 11")
	assert.Equal(t, 6, solution2, "Solution 2 should be 6")
	assert.Nil(t, err, "Error should be nil")
}
