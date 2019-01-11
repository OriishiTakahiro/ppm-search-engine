package store

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"sort"
)

var (
	storedData  = make([]Item, 0, 256)
	jsonFile    = "./data/data.json"
	filePointer *os.File
)

// HitItem represents a item which pair of name and disntance.
type HitItem struct {
	Name     string
	Distance int
}

// Item represents sored item.
type Item interface {
	GetName() string
	GetValues() []int
	FromJSON([]byte) (Item, error)
	DistanceFrom(interface{ GetValues() []int }) int
}

// AddRow adds row to jsons file
func AddRow(row Item) error {

	filePointer, err := os.OpenFile(jsonFile, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	storedData = append(storedData, row)
	jsonBytes, err := json.Marshal(row)
	if err != nil {
		return err
	}
	jsonBytes = append(jsonBytes, '\n')
	filePointer.Write(jsonBytes)

	return nil

}

// OpenAndRead reads data from json file
func OpenAndRead(tmpItem Item) error {

	filePointer, err := os.OpenFile(jsonFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	reader := bufio.NewReaderSize(filePointer, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		item, err := tmpItem.FromJSON(line)
		if err != nil {
			return err
		}
		storedData = append(storedData, item)
	}
	return nil

}

// RowSize adds row to jsons file
func RowSize() int {
	return len(storedData)
}

// SearchNearest adds row to jsons file
func SearchNearest(row Item, limit uint) ([]HitItem, error) {

	candidates := make(map[int][]HitItem)
	for _, v := range storedData {
		d := v.DistanceFrom(row)
		candidates[d] = append(candidates[d], HitItem{Name: v.GetName(), Distance: d})
	}

	results := make([]HitItem, 0, int(limit))

	keys := make([]int, 0, len(storedData))
	for k := range candidates {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		for _, v := range candidates[k] {
			results = append(results, v)
			if len(results) >= int(limit) {
				return results, nil
			}
		}
	}

	return nil, nil
}

// SetJSONPath sets json file path to filename
func SetJSONPath(filename string) {
	jsonFile = filename
}

// Reflesh resets storedData
func Reflesh() {
	storedData = make([]Item, 0, 256)
}
