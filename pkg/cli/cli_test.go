package cli

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	cmd := NewRootCommand("test", "0.0.0")

	out := bytes.NewBufferString("")
	cmd.SetOut(out)
	cmd.SetArgs([]string{})

	assert.Equal(t, "test", cmd.Use, "Use should be test")
	assert.Equal(t, "0.0.0", cmd.Version, "Version should be 0.0.0")

	err := cmd.Execute()

	assert.Nil(t, err, "Error should be nil")
}

func TestNewSubCommandNoRequiredFlags(t *testing.T) {
	var command *cobra.Command
	var arguments []string
	cmd := NewRootCommand("test", "0.0.0")

	NewSubCommand(cmd, "cmd", func(cmd *cobra.Command, args []string) {
		command = cmd
		arguments = args
	})

	out := bytes.NewBufferString("")
	cmd.SetOut(out)
	cmd.SetArgs([]string{"cmd"})

	assert.Equal(t, "test", cmd.Use, "Use should be test")
	assert.Equal(t, "0.0.0", cmd.Version, "Version should be 0.0.0")

	err := cmd.Execute()

	assert.Nil(t, command, "Command should be nil")
	assert.Nil(t, arguments, "Arguments should be nil")
	assert.Contains(t, err.Error(), "file", "Error should contain file")
	assert.Contains(t, err.Error(), "literal", "Error should contain literal")
}

func TestNewSubCommand(t *testing.T) {
	var command *cobra.Command
	var arguments []string
	cmd := NewRootCommand("test", "0.0.0")

	NewSubCommand(cmd, "cmd", func(cmd *cobra.Command, args []string) {
		command = cmd
		arguments = args
	})

	out := bytes.NewBufferString("")
	cmd.SetOut(out)
	cmd.SetArgs([]string{"cmd", "-f", "/tmp/input.txt"})

	assert.Equal(t, "test", cmd.Use, "Use should be test")
	assert.Equal(t, "0.0.0", cmd.Version, "Version should be 0.0.0")

	err := cmd.Execute()
	file, _ := command.Flags().GetString("file")
	literal, _ := command.Flags().GetString("literal")

	assert.Equal(t, "/tmp/input.txt", file, "File flag should be \"/tmp/input.txt\"")
	assert.Equal(t, "", literal, "Literal flag should be empty")
	assert.Len(t, arguments, 0, "Arguments should be empty")
	assert.Nil(t, err, "Error should be nil")
}
