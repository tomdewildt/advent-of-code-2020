package day3

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tomdewildt/advent-of-code-2020/pkg/cli"
	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
)

// AddCommandTo is used to add the command that exectutes the Solve function to
// the main command. This function is also used to load the challenge input. This
// function takes the root command as an input. The result of the Solve function
// will be printed to stdout.
func AddCommandTo(cmd *cobra.Command) {
	cli.NewSubCommand(cmd, "3", func(cmd *cobra.Command, args []string) {
		stream, err := input.FromFlags(cmd.Flags(), "file", "literal")
		if err != nil {
			log.Fatalf("cannot load: %v", err)
		}

		solution1, solution2, err := Solve(stream)
		if err != nil {
			log.Fatalf("cannot solve: %v", err)
		}

		log.Infof("solution (1): %d", solution1)
		log.Infof("solution (2): %d", solution2)
	})
}

// Solve is used to find the solution to the problem. This function takes a stream
// of type io.Reader as input. It returns two integers and nil or 0, 0 and an error
// if one occurred.
func Solve(stream io.Reader) (int, int, error) {
	tileMap, err := input.ToTileMap(stream)
	if err != nil {
		return 0, 0, err
	}

	count1 := traverse(tileMap, 1, 1, '#')
	count2 := traverse(tileMap, 3, 1, '#')
	count3 := traverse(tileMap, 5, 1, '#')
	count4 := traverse(tileMap, 7, 1, '#')
	count5 := traverse(tileMap, 1, 2, '#')

	return count2, (count1 * count2 * count3 * count4 * count5), nil
}

func traverse(tileMap [][]rune, dx int, dy int, char rune) int {
	width := len(tileMap[0])
	height := len(tileMap)
	count := 0
	x := 0
	y := 0

	for y < height {
		if tileMap[y][x%width] == char {
			count++
		}

		x += dx
		y += dy
	}

	return count
}
