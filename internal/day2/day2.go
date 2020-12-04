package day2

import (
	"io"
	"strconv"
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
	cli.NewSubCommand(cmd, "2", func(cmd *cobra.Command, args []string) {
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
	input, err := input.ToStringSlice(stream)
	if err != nil {
		return 0, 0, err
	}

	validCount := 0
	strictlyValidCount := 0
	for _, line := range input {
		policy, password, err := parse(line)
		if err != nil {
			return 0, 0, err
		}

		if policy.isValid(password) {
			validCount++
		}
		if policy.isStrictlyValid(password) {
			strictlyValidCount++
		}
	}

	return validCount, strictlyValidCount, nil
}

type policy struct {
	target rune
	min    int
	max    int
}

func (p policy) isValid(password string) bool {
	count := 0

	for _, letter := range password {
		if letter == p.target {
			count++
		}
	}

	return count >= p.min && count <= p.max
}

func (p policy) isStrictlyValid(password string) bool {
	first := rune(password[p.min-1])
	second := rune(password[p.max-1])

	return (first == p.target) != (second == p.target)
}

func parse(input string) (*policy, string, error) {
	parts := strings.Split(input, " ")

	target := rune(parts[1][0])
	min, err := strconv.Atoi(strings.Split(parts[0], "-")[0])
	if err != nil {
		return nil, "", err
	}
	max, err := strconv.Atoi(strings.Split(parts[0], "-")[1])
	if err != nil {
		return nil, "", err
	}

	return &policy{target: target, min: min, max: max}, parts[2], nil
}
