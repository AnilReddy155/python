//  5-Write a program to sort a list of dictionaries using Lambda.
//  Original list of dictionaries :[{"make": "Nokia", "model": 216, "color": "Black"}, {"make": "Mi Max", "model": "2", "color": "Gold"}, {"make": "Samsung", "model": 7, "color": "Blue"}]

package main

import (
	"log"
	"sort"
)

type Phone struct {
	make  string
	model int
	color string
}

type ColorSorter []Phone

func (a ColorSorter) Len() int           { return len(a) }
func (a ColorSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ColorSorter) Less(i, j int) bool { return a[i].color < a[j].color }

// ModelSorter sorts planets by name.
type ModelSorter []Phone

func (a ModelSorter) Len() int           { return len(a) }
func (a ModelSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ModelSorter) Less(i, j int) bool { return a[i].model < a[j].model }

func main() {
	phones := []Phone{
		{make: "Nokia", model: 216, color: "Black"}, {make: "Mi Max", model: 2, color: "Gold"}, {make: "Samsung", model: 7, color: "Blue"},
	}
	log.Println("unsorted:", phones)

	sort.Sort(ColorSorter(phones))
	log.Println("by color:", phones)

	// sort.Sort(ModelSorter(phones))
	// log.Println("by model:", phones)
}
