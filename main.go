package main

import (
	"fmt"
	"strings"
)

func main() {

	programName := "Handling Employee File"
	fmt.Print("\n")
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("%60s\n", programName)
	fmt.Println(strings.Repeat("-", 100))

	// e := employee.Employee{"1234", "Debaditya", "Bhar", "486 G T Road", "Rajlaxmi Appartment", "Serampore", "WB", "712202"}
	// empData := e.FormatEmpData()
	// emp := employee.Employee{empData[0], empData[1], empData[2], empData[3], empData[4], empData[5], empData[6], empData[7]}
	// emp.ShowEmp()
}
