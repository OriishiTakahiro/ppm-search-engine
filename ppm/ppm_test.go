package ppm

import (
	"encoding/binary"
	"testing"
)

func TestReadPPM(t *testing.T) {
	img, err := ReadPPM(testFile, binary.LittleEndian)
	if err != nil {
		t.Fatalf("Loading test.ppm is failed.\n\t%#v\n", err)
	}
	if img.Width != 145 || img.Height != 144 {
		t.Fatal("Detect uncorrected aspect from the read file.")
	}
	if img.Width*img.Height != len(img.Body) {
		t.Fatal("Data length is illegal.")
	}
}

func TestToHistgram(t *testing.T) {

	img, err := ReadPPM(testFile, binary.LittleEndian)
	row := img.ToHistgram()

	if err != nil {
		t.Fatalf("Converting PPM to Histgram is failed.\n\t%#v\n", err)
	}

	if row.Name != testFile {
		t.Fatalf("Histgram.Name is illegal. \n\tvalue: %s", row.Name)
	}

	rowHasSameLen := len(row.R) == len(row.G) && len(row.R) == len(row.B) && len(row.R) == histgramResolution
	if !rowHasSameLen {
		t.Fatalf("Converting PPM to Histgram is failed. (%s)\n", row.Name)
	}

}
