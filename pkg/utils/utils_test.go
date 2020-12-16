package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPairNotFound(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 10

	result := HasPair(input, target)

	assert.False(t, result, "Result should be false")
}

func TestHasPair(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 50

	result := HasPair(input, target)

	assert.True(t, result, "Result should be true")
}

func TestFindPairNotFound(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 10

	value1, value2, err := FindPair(input, target)

	assert.Equal(t, 0, value1, "Value 1 should be 0")
	assert.Equal(t, 0, value2, "Value 2 should be 0")
	assert.Contains(t, err.Error(), "pair", "Error should contain \"pair\"")
	assert.Contains(t, err.Error(), "not found", "Error should contain \"not found\"")
}

func TestFindPair(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 50

	value1, value2, err := FindPair(input, target)

	assert.Equal(t, 20, value1, "Value 1 should be 20")
	assert.Equal(t, 30, value2, "Value 2 should be 30")
	assert.Nil(t, err, "Error should be nil")
}

func TestHasTripleNotFound(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 10

	result := HasTriple(input, target)

	assert.False(t, result, "Result should be false")
}

func TestHasTriple(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 100

	result := HasTriple(input, target)

	assert.True(t, result, "Result should be true")
}

func TestFindTripleNotFound(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 10

	value1, value2, value3, err := FindTriple(input, target)

	assert.Equal(t, 0, value1, "Value 1 should be 0")
	assert.Equal(t, 0, value2, "Value 2 should be 0")
	assert.Equal(t, 0, value3, "Value 3 should be 0")
	assert.Contains(t, err.Error(), "triple", "Error should contain \"triple\"")
	assert.Contains(t, err.Error(), "not found", "Error should contain \"not found\"")
}

func TestFindTriple(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 100

	value1, value2, value3, err := FindTriple(input, target)

	assert.Equal(t, 40, value1, "Value 1 should be 40")
	assert.Equal(t, 50, value2, "Value 2 should be 50")
	assert.Equal(t, 10, value3, "Value 3 should be 10")
	assert.Nil(t, err, "Error should be nil")
}
