package utils

import (
	"fmt"
)

// HasPair is used to check if a slice of integers contains two numbers that sum
// up to a value. This function takes a slice of integers and target value as
// input. It returns true if the slice contains two numbers that sum up to the
// value.
func HasPair(input []int, target int) bool {
	_, _, err := FindPair(input, target)

	return err == nil
}

// FindPair is used to find two numbers in a slice of integers that sum up to a
// value. This function takes a slice of integers and target value as input. It
// returns two integers and nil or 0, 0 and an error if one occurred.
func FindPair(input []int, target int) (int, int, error) {
	cache := map[int]int{}

	for i, value := range input {
		if i, ok := cache[value]; ok {
			return input[i], value, nil
		}

		cache[target-value] = i
	}

	return 0, 0, fmt.Errorf("pair that sums up to %d not found", target)
}

// HasTriple is used to check if a slice of integers contains three numbers that
// sum up to a value. This function takes a slice of integers and target value as
// input. It returns true if the slice contains three numbers that sum up to the
// value.
func HasTriple(input []int, target int) bool {
	_, _, _, err := FindTriple(input, target)

	return err == nil
}

// FindTriple is used to find three numbers in a slice of integers that sum up to
// a value. This function takes a slice of integers and target value as input. It
// returns three integers and nil or 0, 0 and an error if one occurred.
func FindTriple(input []int, target int) (int, int, int, error) {
	for _, value := range input {
		value1, value2, err := FindPair(input, target-value)

		if err == nil && value1 != value && value2 != value {
			return value1, value2, value, nil
		}
	}

	return 0, 0, 0, fmt.Errorf("triple that sums up to %d not found", target)
}

// FindMinMax is used to find the minimum value and the maximum value in a slice
// of integers. This function takes a slice of integers as input. It returns two
// integers.
func FindMinMax(input []int) (int, int) {
	min, max := 0, 0

	for _, value := range input {
		if value > max {
			max = value
		}
		if min == 0 || value < min {
			min = value
		}
	}

	return min, max
}
