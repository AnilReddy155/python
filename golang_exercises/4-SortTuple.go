package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"English": 88, "Science": 90, "Maths": 97, "Hindi": 80}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		fmt.Println(i, j)
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
}
