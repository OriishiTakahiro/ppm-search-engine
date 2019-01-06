package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ppm-search-engine",
		Short: "Image search search engine.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

// ExecuteCmd : Execute commands.
func ExecuteCmd() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
