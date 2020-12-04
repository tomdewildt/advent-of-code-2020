package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

func TestSolve(t *testing.T) {
	stream := input.FromLiteral("..##...\n#..#..#\n.#...#.\n..#.#.#\n.#.#.##\n#......\n#.##..#\n")

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 3, solution1, "Solution 1 should be 3")
	assert.Equal(t, 36, solution2, "Solution 2 should be 36")
	assert.Nil(t, err, "Error should be nil")
}
