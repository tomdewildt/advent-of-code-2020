package day8

import (
	"fmt"
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
	cli.NewSubCommand(cmd, "8", func(cmd *cobra.Command, args []string) {
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
	code, err := input.ToStringSlice(stream)
	if err != nil {
		return 0, 0, err
	}

	instructions := []instruction{}
	for _, line := range code {
		instruction, err := parse(line)
		if err != nil {
			return 0, 0, err
		}

		instructions = append(instructions, instruction)
	}

	accumulatorHalted, halted, err := process(instructions)
	if !halted && err != nil {
		return 0, 0, err
	}

	accumulatorNotHalted, err := fix(instructions)
	if err != nil {
		return 0, 0, err
	}

	return accumulatorHalted, accumulatorNotHalted, nil
}

type instruction struct {
	operation string
	argument  int
}

func parse(input string) (instruction, error) {
	parts := strings.Split(input, " ")

	operation := parts[0]
	argument, err := strconv.Atoi(parts[1])
	if err != nil {
		return instruction{}, err
	}

	return instruction{operation: operation, argument: argument}, nil
}

func process(instructions []instruction) (int, bool, error) {
	id := 0
	accumulator := 0

	processed := map[int]instruction{}

	for {
		instruction := instructions[id]
		if _, ok := processed[id]; ok {
			return accumulator, true, nil
		}

		processed[id] = instruction

		switch instruction.operation {
		case "acc":
			accumulator += instruction.argument
			id++
		case "jmp":
			id += instruction.argument
		case "nop":
			id++
		default:
			return 0, true, fmt.Errorf("parse: parsing \"%s\": invalid operation", instruction.operation)
		}

		if id == len(instructions) {
			return accumulator, false, nil
		}
	}
}

func fix(instructions []instruction) (int, error) {
	for i, instruction := range instructions {
		original := instruction.operation

		if instruction.operation == "jmp" {
			instructions[i].operation = "nop"
		} else if instruction.operation == "nop" {
			instructions[i].operation = "jmp"
		}

		accumulator, halted, err := process(instructions)
		if err != nil {
			return 0, err
		}

		if !halted {
			return accumulator, nil
		}

		instructions[i].operation = original
	}

	return 0, nil
}
