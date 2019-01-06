package cmd

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/OriishiTakahiro/ppm-search-engine/ppm"
	"github.com/OriishiTakahiro/ppm-search-engine/store"
	"github.com/spf13/cobra"
	"os"
)

type SearchOptions struct {
	limit uint `validate:"min=1,max=100"`
}

var (
	searchOps = &SearchOptions{}

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search similar images from database.",
		Run:   searchFunc,
		Args:  validateSearchArgs,
	}
)

func validateSearchArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a register image file")
	}
	return nil
}

func searchFunc(cmd *cobra.Command, args []string) {

	img, err := ppm.ReadPPM(args[0], binary.LittleEndian)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	hist := img.ToHistgram()
	rows, err := store.SearchTop(hist, searchOps.limit)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for i, row := range rows {
		fmt.Printf("%0d: %s \n\t(Euclid distance %d)\n", i+1, row.Name, row.Distance)
	}

}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().UintVarP(&searchOps.limit, "limit", "l", 10, "A number of showing top-ranking.")
}
