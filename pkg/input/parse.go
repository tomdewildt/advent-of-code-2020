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
	scanner.Split(bufio.ScanWords)

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
