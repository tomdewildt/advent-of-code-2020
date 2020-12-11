package day9

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
	cli.NewSubCommand(cmd, "9", func(cmd *cobra.Command, args []string) {
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

	invalidValue := 0
	for index, value := range input {
		if index >= 25 {
			if _, _, err := findPair(input[index-25:index], value); err != nil {
				invalidValue = value
			}
		}
	}

	weaknessValue := 0
	for i := 0; i < len(input)-1; i++ {
		sum := input[i]
		for j := i + 1; j < len(input); j++ {
			sum += input[j]
			if sum == invalidValue {
				min, max := findMinMax(input[i : j+1])
				weaknessValue = min + max
				break
			} else if sum > invalidValue {
				break
			}
		}
	}

	return invalidValue, weaknessValue, nil
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

func findMinMax(input []int) (int, int) {
	min, max := 0, 0

	for _, value := range input {
		if value > max {
			max = value
		}
		if min == 0 || value < min {
			min = value
		}
	}

	return min, max
}
