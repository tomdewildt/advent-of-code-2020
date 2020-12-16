package input

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestFromFlagsInvalidFileFlag(t *testing.T) {
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("file", "", "test input file")
	flags.String("literal", "", "test input literal")

	stream, err := FromFlags(flags, "test", "literal")

	assert.Nil(t, stream, "Stream should be nil")
	assert.Contains(t, err.Error(), "test", "Error should contain \"test\"")
}

func TestFromFlagsInvalidLiteralFlag(t *testing.T) {
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("file", "/tmp/input.txt", "test input file")
	flags.String("literal", "", "test input literal")

	stream, err := FromFlags(flags, "file", "test")

	assert.Nil(t, stream, "Stream should be nil")
	assert.Contains(t, err.Error(), "test", "Error should contain \"test\"")
}

func TestFromFlagsMissingFlags(t *testing.T) {
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("file", "", "test input file")
	flags.String("literal", "", "test input literal")

	stream, err := FromFlags(flags, "file", "literal")

	assert.Nil(t, stream, "Stream should be nil")
	assert.Contains(t, err.Error(), "file", "Error should contain \"file\"")
	assert.Contains(t, err.Error(), "literal", "Error should contain \"literal\"")
}

func TestFromFlagsFromFile(t *testing.T) {
	path := "/tmp/input.txt"
	os.Create(path)
	defer func() { os.Remove(path) }()

	flags := pflag.NewFlagSet("test", pflag.PanicOnError)
	flags.String("file", "/tmp/input.txt", "test input file")
	flags.String("literal", "", "test input literal")

	stream, err := FromFlags(flags, "file", "literal")
	data, _ := ioutil.ReadAll(stream)

	assert.Equal(t, "", string(data), "Stream should be empty")
	assert.Nil(t, err, "Error should be nil")
}

func TestFromFlagsFromLiteral(t *testing.T) {
	flags := pflag.NewFlagSet("test", pflag.PanicOnError)
	flags.String("file", "", "test input file")
	flags.String("literal", "literal", "test input literal")

	stream, err := FromFlags(flags, "file", "literal")
	data, _ := ioutil.ReadAll(stream)

	assert.Equal(t, "literal", string(data), "Stream should be \"literal\"")
	assert.Nil(t, err, "Error should be nil")
}

func TestFromFile(t *testing.T) {
	path := "/tmp/input.txt"
	os.Create(path)
	defer func() { os.Remove(path) }()

	stream, err := FromFile(path)
	data, _ := ioutil.ReadAll(stream)

	assert.Equal(t, "", string(data), "Stream should be empty")
	assert.Nil(t, err, "Error should be nil")
}

func TestFromLiteral(t *testing.T) {
	text := "test"

	stream := FromLiteral(text)
	data, _ := ioutil.ReadAll(stream)

	assert.Equalf(t, text, string(data), "Stream should be \"%s\"", text)
}
