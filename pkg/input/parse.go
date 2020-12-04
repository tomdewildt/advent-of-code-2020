package input

import (
	"bufio"
	"io"
	"strconv"
)

// ToIntSlice is used to convert an file into a slice of integers. This function
// takes an stream of type io.Reader as input. It returns an slice of integers
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

// ToStringSlice is used to convert an file into a slice of strings. This function
// takes an stream of type io.Reader as input. It returns an slice of strings and
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

// ToTileMap is used to convert an file into a 2d slice of runes. This function takes
// an stream of type io.Reader as input. It returns an 2d slice of runes and nil or
// nil and an error if one occurred.
func ToTileMap(stream io.Reader) ([][]rune, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	var result [][]rune

	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}

	return result, scanner.Err()
}
