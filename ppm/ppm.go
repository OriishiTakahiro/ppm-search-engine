package ppm

import (
	"encoding/binary"
	"io"
	"os"
	sc "strconv"
)

// PPM represents .ppm binary file.
type PPM struct {
	Name   string
	Width  int
	Height int
	Body   []RGB
}

// RGB represents a triple of RGB colors.
type RGB struct {
	R byte
	G byte
	B byte
}

// NewPPM instantiates PPM objects with initial values.
func NewPPM(filename string) PPM {
	obj := PPM{
		Name:   filename,
		Width:  0,
		Height: 0,
		Body:   make([]RGB, 0),
	}
	return obj
}

// ToHistgram converts PPM to Histgram.
func (c PPM) ToHistgram() (result Histgram) {
	result = NewHistgram(c.Name)
	for _, v := range c.Body {
		result.R[v.R/histgramResolution]++
		result.G[v.G/histgramResolution]++
		result.B[v.B/histgramResolution]++
	}
	return
}

// ReadPPM loads specified a .ppm file and convert to PPM object.
func ReadPPM(filename string, order binary.ByteOrder) (PPM, error) {

	result := NewPPM(filename)
	file, err := os.Open(filename)
	if err != nil {
		return PPM{}, err
	}
	defer file.Close()

	// Read header values.
	v := make([]byte, 3)
	file.Seek(3, 0)
	binary.Read(file, order, &v)
	result.Width, _ = sc.Atoi(string(v))

	file.Seek(7, 0)
	binary.Read(file, order, &v)
	result.Height, _ = sc.Atoi(string(v))

	// Read body values.
	file.Seek(ppmHeaderSize, 0) // skip header
	var val RGB
	for {
		if err := binary.Read(file, order, &val); err == io.EOF {
			break
		}
		result.Body = append(result.Body, val)
	}

	return result, nil
}
