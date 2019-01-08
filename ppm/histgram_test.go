package ppm

import (
	"encoding/binary"
	"testing"
)

func TestEuclidDistance(t *testing.T) {

	img, _ := ReadPPM(testFile, binary.LittleEndian)
	row := img.ToHistgram()
	row2 := img.ToHistgram()

	if EuclidDistance(row, row2) != 0 {
		t.Fatal("Different distance from 2 same images.")
	}

	row2.R[3] = 0
	row2.G[2] = 0
	if EuclidDistance(row, row2) == 0 {
		t.Fatal("Same distance from 2 different images.")
	}
}
