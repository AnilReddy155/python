// 7-Write a program to access only unique key value of a Python object.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// data in JSON format which
	// is to be decoded
	Data := []byte(`{
		"Name": "India",  
		"Capital": "New Delhi",
		"Continent": "Asia",
		"Name": "Bharat"
	}`)

	country1 := make(map[string]json.RawMessage)

	// decoding country1 struct
	// from json format
	err := json.Unmarshal(Data, &country1)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// a string slice to hold the keys
	k := make([]string, len(country1))

	// iteration counter
	i := 0

	// printing details of keys into k
	for s, _ := range country1 {
		k[i] = s
		i++
	}
	fmt.Printf("%#v\n", k)
}
