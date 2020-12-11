package day10

import (
	"io"
	"sort"

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
	cli.NewSubCommand(cmd, "10", func(cmd *cobra.Command, args []string) {
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
	sort.Ints(input)

	diff1 := 0
	diff3 := 1
	prev := 0
	for _, value := range input {
		delta := value - prev
		if delta == 3 {
			diff3++
		} else if delta == 1 {
			diff1++
		}

		prev = value
	}

	input = append([]int{0}, input...)
	permutations := permute(input, map[int]int{})

	return diff1 * diff3, permutations, nil
}

func permute(input []int, lookup map[int]int) int {
	if len(input) == 1 {
		return 1
	}

	if count, ok := lookup[input[0]]; ok {
		return count
	}

	count := 0
	for i := 1; i < len(input); i++ {
		if input[i]-input[0] <= 3 {
			count += permute(input[i:], lookup)
		}
	}
	lookup[input[0]] = count

	return count
}
