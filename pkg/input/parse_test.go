package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntSliceInvalidStream(t *testing.T) {
	stream := FromLiteral("a\n2\n3\n")

	result, err := ToIntSlice(stream)

	assert.Nil(t, result, "Result should be nil")
	assert.Contains(t, err.Error(), "invalid syntax", "Error should contain \"invalid syntax\"")
}

func TestToIntSlice(t *testing.T) {
	stream := FromLiteral("1\n2\n3\n")

	result, err := ToIntSlice(stream)

	assert.Equal(t, []int{1, 2, 3}, result, "Result should be [1,2,3]")
	assert.Nil(t, err, "Error should be nil")
}

func TestToStringSlice(t *testing.T) {
	stream := FromLiteral("a\nb\nc\n")

	result, err := ToStringSlice(stream)

	assert.Equal(t, []string{"a", "b", "c"}, result, "Result should be [a,b,c]")
	assert.Nil(t, err, "Error should be nil")
}

func TestToGroupedStringSlice(t *testing.T) {
	stream := FromLiteral("abc\ndef\n\nghi\njkl\n\nmno")

	result, err := ToGroupedStringSlice(stream)

	assert.Equal(t, []string{"abc def", "ghi jkl", "mno"}, result, "Result should be [abc def,ghi jkl,mno]")
	assert.Nil(t, err, "Error should be nil")
}

func TestToTileMap(t *testing.T) {
	stream := FromLiteral(".#.\n#.#\n.#.\n")

	result, err := ToTileMap(stream)

	assert.Equal(t, [][]rune{{'.', '#', '.'}, {'#', '.', '#'}, {'.', '#', '.'}}, result, "Result should be [[.,#,.],[#,.,#],[.,#,.]]")
	assert.Nil(t, err, "Error should be nil")
}
