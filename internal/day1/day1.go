package day1

import (
	"errors"
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
	cli.NewSubCommand(cmd, "1", func(cmd *cobra.Command, args []string) {
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

// Solve is used to find the solution to the problem. This function takes an stream
// of type io.Reader as input. It returns two integers and nil or 0, 0 and an error
// if one occurred.
func Solve(stream io.Reader) (int, int, error) {
	input, err := input.ToIntSlice(stream)
	if err != nil {
		return 0, 0, err
	}

	pair1, pair2, err := findPair(input, 2020)
	if err != nil {
		return 0, 0, err
	}

	triple1, triple2, triple3, err := findTriple(input, 2020)
	if err != nil {
		return 0, 0, err
	}

	return (pair1 * pair2), (triple1 * triple2 * triple3), nil
}

func findPair(input []int, target int) (int, int, error) {
	cache := map[int]int{}

	for index, value := range input {
		if index, ok := cache[value]; ok {
			return input[index], value, nil
		}

		cache[target-value] = index
	}

	return 0, 0, errors.New("pair not found")
}

func findTriple(input []int, target int) (int, int, int, error) {
	for _, value := range input {
		value1, value2, err := findPair(input, target-value)

		if err == nil && value1 != value && value2 != value {
			return value1, value2, value, nil
		}
	}

	return 0, 0, 0, errors.New("triple not found")
}
