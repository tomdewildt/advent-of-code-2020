package main

import (
	"os"

	"github.com/tomdewildt/advent-of-code-2020/internal/day1"
	"github.com/tomdewildt/advent-of-code-2020/pkg/cli"
)

// Name of the binary
var Name string

// Version of the binary
var Version string

func main() {
	cmd := cli.NewRootCommand(Name, Version)

	day1.AddCommandTo(cmd)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
