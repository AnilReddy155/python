// 8-Write a program to read each row from a given csv file and print a list of strings

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	id   int
	name string
	age  int
}

func main() {
	// read data from CSV file

	csvFile, err := os.Open("../../read.csv")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(csvFile)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	fmt.Println(csvData)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Employee
	var allRecords []Employee

	for i, each := range csvData {
		if i >= 1 {
			eachRec := strings.Split(each[0], ",")
			id := eachRec[0]
			age := eachRec[2]
			oneRecord.id, _ = strconv.Atoi(strings.Trim(id, " "))
			oneRecord.name = string(eachRec[1])
			oneRecord.age, _ = strconv.Atoi(strings.Trim(age, " "))
			// for _, rec := range eachRec {
			// 	println(eachRec[0])

			// 	//println(rec)
			// 	// oneRecord.id = strconv.Atoi(string(rec[0]))
			// 	// oneRecord.name = string(rec[1])
			// 	// oneRecord.age = strconv.Atoi(string(rec[2]))
			// }
			allRecords = append(allRecords, oneRecord)
		}

	}

	fmt.Println(allRecords)

}
