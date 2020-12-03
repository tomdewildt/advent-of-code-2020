package input

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

// FromFlags is used to determine how a challenge input should be loaded. If the
// file flag is set the FromFile function will be used, if the literal flag is set
// the FromLiteral function will be used. This function takes a flagset of type
// pflag.FlagSet, a file key and an literal key as inputs. It returns an stream of
// type io.Reader and nil or nil and an error if one occurred.
func FromFlags(flags *pflag.FlagSet, fileKey string, literalKey string) (io.Reader, error) {
	file, err := flags.GetString(fileKey)
	if err != nil {
		return nil, err
	}

	literal, err := flags.GetString(literalKey)
	if err != nil {
		return nil, err
	}

	if file != "" {
		return FromFile(file)
	}
	return FromLiteral(literal), nil
}

// FromFile is used to load a challenge input from a file. This function takes the
// filename as input. It returns an stream of type io.Reader and nil or nil and an
// error if one occurred.
func FromFile(path string) (io.Reader, error) {
	return os.Open(path)
}

// FromLiteral is used to load a challenge input from a string. This function takes
// the text as input. It returns an stream of type io.Reader and nil or nil and an
// error if one occurred.
func FromLiteral(text string) io.Reader {
	return strings.NewReader(text)
}
