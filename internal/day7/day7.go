package day7

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
	cli.NewSubCommand(cmd, "7", func(cmd *cobra.Command, args []string) {
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

	bags := map[bag]map[bag]int{}
	for _, line := range input {
		bag, contains, err := parse(line)
		if err != nil {
			return 0, 0, err
		}

		bags[bag] = contains
	}

	containsCount := 0
	requiredCount := 0
	for current := range bags {
		if contains(bags, current, bag{modifier: "shiny", color: "gold"}) {
			containsCount++
		}
	}

	requiredCount = required(bags, bag{modifier: "shiny", color: "gold"}) - 1

	return containsCount, requiredCount, nil
}

type bag struct {
	modifier string
	color    string
}

func parse(input string) (bag, map[bag]int, error) {
	parts := strings.Split(input, " bags contain ")

	contains := map[bag]int{}
	if parts[1] != "no other bags." {
		for _, line := range strings.Split(parts[1], ", ") {
			parts := strings.Split(line, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return bag{}, contains, err
			}

			contains[bag{modifier: parts[1], color: parts[2]}] = count
		}
	}

	parts = strings.Split(parts[0], " ")
	bag := bag{
		modifier: parts[0],
		color:    parts[1],
	}

	return bag, contains, nil
}

func contains(bags map[bag]map[bag]int, start bag, target bag) bool {
	if count := bags[start][target]; count > 0 {
		return true
	}

	for bag := range bags[start] {
		if contains(bags, bag, target) {
			return true
		}
	}

	return false
}

func required(bags map[bag]map[bag]int, target bag) int {
	total := 1

	for bag, count := range bags[target] {
		total += count * required(bags, bag)
	}

	return total
}
