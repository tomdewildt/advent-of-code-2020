package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

const tileMap = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestSolve(t *testing.T) {
	stream := input.FromLiteral(tileMap)

	solution1, solution2, err := Solve(stream)

	assert.Equal(t, 7, solution1, "Solution 1 should be 7")
	assert.Equal(t, 336, solution2, "Solution 2 should be 336")
	assert.Nil(t, err, "Error should be nil")
}
