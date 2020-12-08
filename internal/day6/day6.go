package day6

import (
	"io"
	"strings"

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
	cli.NewSubCommand(cmd, "6", func(cmd *cobra.Command, args []string) {
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
	input, err := input.ToGroupedStringSlice(stream)
	if err != nil {
		return 0, 0, err
	}

	answeredYes := 0
	answeredAllYes := 0
	for _, group := range input {
		answers := map[rune]int{}

		for _, answer := range group {
			if answer != ' ' {
				answers[answer]++
			}
		}
		answeredYes += len(answers)

		size := strings.Count(group, " ") + 1
		for _, count := range answers {
			if count == size {
				answeredAllYes++
			}
		}
	}

	return answeredYes, answeredAllYes, nil
}
