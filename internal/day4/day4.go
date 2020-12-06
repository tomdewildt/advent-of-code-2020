package day4

import (
	"io"
	"regexp"
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
	cli.NewSubCommand(cmd, "4", func(cmd *cobra.Command, args []string) {
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
	currentPassport := passport{}
	for _, line := range input {
		currentPassport.parse(line)

		if line == "" || line == input[len(input)-1] {
			if currentPassport.isValid() {
				validCount++
			}
			if currentPassport.isStrictlyValid() {
				strictlyValidCount++
			}

			currentPassport = passport{}
		}
	}

	return validCount, strictlyValidCount, nil
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) isValid() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p passport) isStrictlyValid() bool {
	if byr, err := strconv.Atoi(p.byr); err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	if iyr, err := strconv.Atoi(p.iyr); err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	if eyr, err := strconv.Atoi(p.eyr); err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	if strings.HasSuffix(p.hgt, "cm") {
		if hgt, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm")); err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		if hgt, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in")); err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	if !regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(p.hcl) {
		return false
	}
	if !regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$").MatchString(p.ecl) {
		return false
	}
	if _, err := strconv.Atoi(p.pid); err != nil || len(p.pid) != 9 {
		return false
	}

	return true
}

func (p *passport) parse(input string) {
	fields := strings.Split(input, " ")

	for _, field := range fields {
		pairs := strings.Split(field, ":")

		switch pairs[0] {
		case "byr":
			p.byr = pairs[1]
		case "iyr":
			p.iyr = pairs[1]
		case "eyr":
			p.eyr = pairs[1]
		case "hgt":
			p.hgt = pairs[1]
		case "hcl":
			p.hcl = pairs[1]
		case "ecl":
			p.ecl = pairs[1]
		case "pid":
			p.pid = pairs[1]
		case "cid":
			p.cid = pairs[1]
		}
	}
}
