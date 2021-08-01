package cmd

import (
	"context"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

// NewTreeCommand create cobra command for Tree
func NewTreeCommand(ctx context.Context) *cobra.Command {
	opts := newTreeOptions()
	var treeCmd = &cobra.Command{
		Use:   "tree",
		Short: "create a family tree",
		Long:  "Tree will be a diagram of the family tree",
		Run: func(cmd *cobra.Command, args []string) {
			printVersionInfo()

			if err := opts.loadFamilyFromFile(); err != nil {
				opts.Logger.Fatalf("failed to load family from tree file: %v", err)
			}

			if err := opts.validate(); err != nil {
				opts.Logger.Fatalf("failed to validate family from tree file: %v", err)
			}

			opts.complete()

			optsJSON, err := yaml.Marshal(opts.Config)
			if err != nil {
				opts.Logger.Fatalf("failed to marshal the options: %v", err)
			}
			opts.Logger.Infof("%s", optsJSON)

			if err := opts.run(ctx); err != nil {
				opts.Logger.Fatalf("failed to run family tree: %v", err)
			}
		},
	}

	opts.addFlags(treeCmd.Flags())
	return treeCmd
}
