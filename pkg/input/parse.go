package input

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ToInt is used to convert a stream into an integer. This function takes a stream
// of type io.Reader as input. It returns an integer and nil or 0 and an error if
// one occurred.
func ToInt(stream io.Reader) (int, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		result = value
	}

	return result, nil
}

// ToString is used to convert a stream into a string. This function takes a
// stream of type io.Reader as input. It returns a string and nil or an empty
// string and an error if one occurred.
func ToString(stream io.Reader) (string, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result string

	for scanner.Scan() {
		result = scanner.Text()
	}

	return result, nil
}

// ToIntSlice is used to convert a stream into a slice of integers. This function
// takes a stream of type io.Reader as input. It returns a slice of integers
// and nil or nil and an error if one occurred.
func ToIntSlice(stream io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result []int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, scanner.Err()
}

// ToStringSlice is used to convert a stream into a slice of strings. This function
// takes a stream of type io.Reader as input. It returns a slice of strings and
// nil or nil and an error if one occurred.
func ToStringSlice(stream io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

// ToGroupedStringSlice is used to convert a stream into a slice of grouped strings.
// The groups must be separated by empty lines. This function takes a stream of type
// io.Reader as input. It returns a slice of grouped strings and nil or nil and an
// error if one occurred.
func ToGroupedStringSlice(stream io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result []string
	var group string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			result = append(result, group)
			group = ""
		} else {
			group = strings.TrimSpace(group + " " + line)
		}
	}

	if group != "" {
		result = append(result, group)
	}

	return result, scanner.Err()
}

// ToTileMap is used to convert a stream into a 2d slice of runes. This function
// takes a stream of type io.Reader as input. It returns a 2d slice of runes and
// nil or nil and an error if one occurred.
func ToTileMap(stream io.Reader) ([][]rune, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result [][]rune

	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}

	return result, scanner.Err()
}
