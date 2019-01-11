package ppm

import (
	"encoding/json"
	"github.com/OriishiTakahiro/ppm-search-engine/store"
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

// FromJSON instantiates Histgram object from json data.
func (h Histgram) FromJSON(jsonData []byte) (store.Item, error) {
	result := Histgram{}
	err := json.Unmarshal(jsonData, &result)
	return result, err
}

// DistanceFrom calculate euclid distance between lhs and rhs.
func (lhs Histgram) DistanceFrom(rhs interface{ GetValues() []int }) int {

	d := 0.0
	lhsValues := lhs.GetValues()
	rhsValues := rhs.GetValues()

	for i, v := range lhsValues {
		d += math.Abs(math.Pow(float64(v), 2) - math.Pow(float64(rhsValues[i]), 2))
	}

	return int(d)
}

// GetName archive Name of a histgram item.
func (recv Histgram) GetName() string {
	return recv.Name
}

// GetValues returns a series of R, G, B histgram
func (recv Histgram) GetValues() []int {
	result := append(recv.R, recv.G...)
	result = append(result, recv.B...)
	return result
}
