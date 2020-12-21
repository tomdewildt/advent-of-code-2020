package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

const numbers = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestSolveInvalidNumber(t *testing.T) {
	stream := input.FromLiteral("abc")
	preamble := 5

	solution1, solution2, err := Solve(stream, preamble)

	assert.Equal(t, 0, solution1, "Solution 1 should be 0")
	assert.Equal(t, 0, solution2, "Solution 2 should be 0")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestSolve(t *testing.T) {
	stream := input.FromLiteral(numbers)
	preamble := 5

	solution1, solution2, err := Solve(stream, preamble)

	assert.Equal(t, 127, solution1, "Solution 1 should be 127")
	assert.Equal(t, 62, solution2, "Solution 2 should be 62")
	assert.Nil(t, err, "Error should be nil")
}
