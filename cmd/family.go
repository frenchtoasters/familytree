package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// NewFamilyTreeCommand is the base command for the family tree app
func NewFamilyCommand(ctx context.Context) *cobra.Command {
	var RootCmd = &cobra.Command{
		Use:   "family",
		Short: "command line utility for creating family trees",
		Long:  "This command line utility is supporting the Family Tree web app",
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				printVersionInfo()
			}
		},
	}

	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "print version info")
	RootCmd.AddCommand(NewTreeCommand(ctx))

	return RootCmd
}
