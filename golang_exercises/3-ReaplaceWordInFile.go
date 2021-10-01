package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	read, err := ioutil.ReadFile("word-replace.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(read))
	//fmt.Println(path)
	data := string(read)
	fmt.Println(strings.Count(data, "pyton"))
	fmt.Println(strings.Count(data, "ga"))
	newContents := strings.Replace(data, "pyton", "python", -1)

	newContents = strings.Replace(string(newContents), "ga", "go", -1)

	//fmt.Println(newContents)

	err = ioutil.WriteFile("word-replace.txt", []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
