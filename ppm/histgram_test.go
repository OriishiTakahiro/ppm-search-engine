package ppm

import (
	"encoding/binary"
	"testing"
)

func TestDistanceFrom(t *testing.T) {

	img, _ := ReadPPM(testFile, binary.LittleEndian)
	row := img.ToHistgram()
	row2 := img.ToHistgram()

	if row.DistanceFrom(row2) != 0 {
		t.Fatal("Different distance from 2 same images.")
	}

	row2.R[3] = 0
	row2.G[2] = 0
	if row.DistanceFrom(row2) == 0 {
		t.Fatal("Same distance from 2 different images.")
	}
}

func TestGetValues(t *testing.T) {
	img, _ := ReadPPM(testFile, binary.LittleEndian)

	hist := img.ToHistgram()
	values := hist.GetValues()
	sumLen := len(hist.R) + len(hist.G) + len(hist.B)

	if sumLen != len(values) || sumLen != histgramResolution*3 {
		t.Fatalf("Detect illegal length of Histgram.GetValues() = %d, but must be %d.\n", len(values), histgramResolution*3)
	}

	// order of values is R, G, B
	for i, v := range values[0 : histgramResolution-1] {
		if v != hist.R[i] {
			t.Fatalf("Different correspond value %d != %d", v, hist.R[i])
		}
	}
	for i, v := range values[histgramResolution : 2*histgramResolution-1] {
		if v != hist.G[i] {
			t.Fatalf("Different correspond value %d != %d", v, hist.G[i])
		}
	}
	for i, v := range values[2*histgramResolution:] {
		if v != hist.B[i] {
			t.Fatalf("Different correspond value %d != %d", v, hist.B[i])
		}
	}
}

func TestGetName(t *testing.T) {
	img, _ := ReadPPM(testFile, binary.LittleEndian)
	if img.ToHistgram().GetName() != testFile {
		t.Fatal("Filename is not registered.")
	}
}
