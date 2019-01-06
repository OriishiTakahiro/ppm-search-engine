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

var (
	registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register images to database.",
		Run:   registerFunc,
		Args:  validateRegisterArgs,
	}
)

func registerFunc(cmd *cobra.Command, args []string) {

	img, err := ppm.ReadPPM(args[0], binary.LittleEndian)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	row := img.ToHistgram()
	if err := store.AddRow(row); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("Inputed image is registerd!\n%s", args[0])
}

func validateRegisterArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a register image file")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
