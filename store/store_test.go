package store

import (
	"math"
	"testing"
)

type hoge struct {
	name string
	vals []int
}

func (h hoge) GetName() string                 { return h.name }
func (h hoge) GetValues() []int                { return h.vals }
func (h hoge) FromJSON(b []byte) (Item, error) { return nil, nil }
func (h hoge) DistanceFrom(r interface{ GetValues() []int }) int {
	d := 0.0
	hv := h.GetValues()
	rv := r.GetValues()
	for i, v := range hv {
		d += math.Abs(float64(v) - float64(rv[i]))
	}
	return int(d)
}

func TestAddRow(t *testing.T) {

	SetJSONPath("/dev/null")

	Reflesh()

	before := RowSize()
	AddRow(hoge{})
	if RowSize() != before+1 {
		t.Fatal("AddRow cannot increment size")
	}

}

func TestSearchNearest(t *testing.T) {

	list := []hoge{
		hoge{vals: []int{0, 8, 0, 0}, name: "h0"},
		hoge{vals: []int{0, 0, 10, 5}, name: "h1"},
		hoge{vals: []int{0, 0, 0, 0}, name: "h2"},
		hoge{vals: []int{0, 0, 10, 0}, name: "h3"},
		hoge{vals: []int{0, 2, 0, 1}, name: "h4"},
	}
	target := hoge{vals: []int{0, 0, 0, 0}, name: "target"}

	SetJSONPath("/dev/null")
	Reflesh()

	for _, v := range list {
		AddRow(v)
	}

	result, _ := SearchNearest(target, 3)
	if len(result) != 3 {
		t.Fatalf("Illegal length of search result %d.", len(result))
	}
	if result[0].Name != list[2].name {
		t.Fatalf("Illegal top of search result (%s, %s).", result[0].Name, list[2].name)
	}
	if result[2].Name != list[0].name {
		t.Fatalf("Illegal 3rd of search result (%s, %s).", result[2].Name, list[0].name)
	}

}
