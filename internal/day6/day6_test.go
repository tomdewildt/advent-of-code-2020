package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("abc\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 8, solution1, "Solution 1 should be 8")
	assert.Equal(t, 3, solution2, "Solution 2 should be 3")
	assert.Nil(t, err, "Error should be nil")
}
