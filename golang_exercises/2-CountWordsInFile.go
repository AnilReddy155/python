// 2- Search and count the word available in file.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	words := []string{}
	count := 0
	read, err := ioutil.ReadFile("word-count.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(read)

	words = strings.Split(data, " ")
	count = len(words)
	fmt.Println(count)

	// for _, word := range words {
	// 	fmt.Println(word)
	// 	count = count + 1
	// }
	// fmt.Println(count)
	fmt.Println(" Number of Words in Given File is : ", count)

}
