package cmd

import (
	"fmt"
	"github.com/OriishiTakahiro/ppm-search-engine/ppm"
	"github.com/OriishiTakahiro/ppm-search-engine/store"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ppm-search-engine",
		Short: "Image search search engine.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	if err := store.OpenAndRead(ppm.Histgram{}); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// ExecuteCmd : Execute commands.
func ExecuteCmd() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
