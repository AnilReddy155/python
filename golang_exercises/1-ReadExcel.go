// 1-Create employee data excel file with duplicate record  that contain the fallowing field
// Emp id, Emp Name, Emp Salary
// Write a python code to display duplicate employee record and count no of duplicate record available into a employee data file.

package main

import (
	//"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
	//"github.com/tealeg/xlsx"
)

type emp struct {
	empId     int
	empName   string
	empSalary int
}

func main() {
	// empployees := createEmployees()

	f, err := excelize.OpenFile("Book.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	//fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
		return
	}
	employees := []emp{}
	var e emp
	for i, row := range rows {

		if i > 0 {
			e.empId, _ = strconv.Atoi(row[0])
			e.empName = row[1]
			e.empSalary, _ = strconv.Atoi(row[2])
			employees = append(employees, e)
		}

	}
	count := 1
	a := 0
	fmt.Println(a == 0)
	em := make(map[emp]int)

	for _, entry := range employees {
		if value := em[entry]; value == 0 {
			em[entry] = count
		} else {
			em[entry] = em[entry] + 1
		}
	}

	fmt.Println(em)
}

func createEmployees() []emp {
	employees := []emp{
		emp{
			empId:     1,
			empName:   "anil",
			empSalary: 1000,
		},
		emp{
			empId:     1,
			empName:   "anil",
			empSalary: 1000,
		}, emp{
			empId:     3,
			empName:   "suneel",
			empSalary: 1000,
		},
		emp{
			empId:     1,
			empName:   "anil",
			empSalary: 1000,
		},
		emp{
			empId:     1,
			empName:   "anil",
			empSalary: 1000,
		}, emp{
			empId:     3,
			empName:   "suneel",
			empSalary: 1000,
		},
	}

	return employees
}
