package day9

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tomdewildt/advent-of-code-2020/pkg/cli"
	"github.com/tomdewildt/advent-of-code-2020/pkg/input"
	"github.com/tomdewildt/advent-of-code-2020/pkg/utils"
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

		solution1, solution2, err := Solve(stream, 25)
		if err != nil {
			log.Fatalf("cannot solve: %v", err)
		}

		log.Infof("solution (1): %d", solution1)
		log.Infof("solution (2): %d", solution2)
	})
}

// Solve is used to find the solution to the problem. This function takes a stream
// of type io.Reader and the preamble as input. It returns two integers and nil or
// 0, 0 and an error if one occurred.
func Solve(stream io.Reader, preamble int) (int, int, error) {
	numbers, err := input.ToIntSlice(stream)
	if err != nil {
		return 0, 0, err
	}

	invalidValue := 0
	for i, number := range numbers {
		if i >= preamble {
			if ok := utils.HasPair(numbers[i-preamble:i], number); !ok {
				invalidValue = number
			}
		}
	}

	weaknessValue := 0
	for i := 0; i < len(numbers)-1; i++ {
		sum := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if sum == invalidValue {
				min, max := utils.FindMinMax(numbers[i : j+1])
				weaknessValue = min + max
				break
			} else if sum > invalidValue {
				break
			}
		}
	}

	return invalidValue, weaknessValue, nil
}
