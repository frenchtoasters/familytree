package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// NewFamilyTreeCommand is the base command for the family tree app
func NewFamilyTreeCommand(ctx context.Context) *cobra.Command {
	var RootCmd = &cobra.Command{
		Use:   "familytree",
		Short: "command line utility for creating family trees",
		Long:  "This command line utility is supporting the Family Tree web app",
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				printVersionInfo()
			}
		},
	}

	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "print version info")
	RootCmd.AddCommand(NewFamilyTreeCommand(ctx))

	return RootCmd
}
