package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NewRootCommand is used to create a new instance of a cobra.Command. This function
// takes a name and version as input. It returns a command of type cobra.Command.
func NewRootCommand(name string, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     name,
		Short:   "Advent of code 2020 solutions.",
		Long:    "Solutions for the advent of code 2020 problems.",
		Version: version,
	}

	cmd.SetHelpCommand(&cobra.Command{Hidden: true})

	return cmd
}

// NewSubCommand is used to add a new sub command to a cobra.Command. This function
// takes a root command, name and run function as input. The run function will be
// invoked when the command is executed and all the flags pass validation.
func NewSubCommand(root *cobra.Command, name string, run func(cmd *cobra.Command, args []string)) {
	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Day %s solution.", name),
		Long:  fmt.Sprintf("Solution for the advent of code 2020 day %s problem.", name),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			file, _ := cmd.Flags().GetString("file")
			literal, _ := cmd.Flags().GetString("literal")

			if file == "" && literal == "" {
				return errors.New("required flag \"file\" or \"literal\" not set")
			}

			return nil
		},
		Run: run,
	}

	cmd.PersistentFlags().StringP("file", "f", "", "challenge input file")
	cmd.PersistentFlags().StringP("literal", "l", "", "challenge input literal")

	root.AddCommand(cmd)
}
