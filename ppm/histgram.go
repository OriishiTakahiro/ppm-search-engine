package ppm

import (
	"math"
)

// Histgram represents RGB-Histgram.
type Histgram struct {
	Name string `json:"name"`
	R    []int  `json:"R"`
	G    []int  `json:"G"`
	B    []int  `json:"B"`
}

// NewHistgram instantiates Histgram object with initial values.
func NewHistgram(filename string) (obj Histgram) {
	obj = Histgram{
		Name: filename,
		R:    make([]int, 16),
		G:    make([]int, 16),
		B:    make([]int, 16),
	}
	return
}

// EuclidDistance calculate euclid distance between lhs and rhs.
func EuclidDistance(lhs Histgram, rhs Histgram) (d int) {

	d = 0
	tmpR := 0.0
	tmpG := 0.0
	tmpB := 0.0
	for i, v := range lhs.R {
		tmpR += math.Pow(float64(v-rhs.R[i]), 2.0)
	}
	for i, v := range lhs.G {
		tmpG += math.Pow(float64(v-rhs.G[i]), 2.0)
	}
	for i, v := range lhs.B {
		tmpB += math.Pow(float64(v-rhs.B[i]), 2.0)
	}
	tmpR = tmpR / histgramResolution
	tmpG = tmpG / histgramResolution
	tmpB = tmpB / histgramResolution

	d = int((tmpR + tmpG + tmpB) / 3.0)

	return
}
