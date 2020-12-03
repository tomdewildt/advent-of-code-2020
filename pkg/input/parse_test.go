package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntSliceInvalidStream(t *testing.T) {
	stream := FromLiteral("a\n2\n3\n")

	result, err := ToIntSlice(stream)

	assert.Nil(t, result, "Result should be nil")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain invalid syntax")
}

func TestToIntSlice(t *testing.T) {
	stream := FromLiteral("1\n2\n3\n")

	result, err := ToIntSlice(stream)

	assert.Equal(t, []int{1, 2, 3}, result, "Result should be [1,2,3]")
	assert.Nil(t, err, "Error should be nil")
}