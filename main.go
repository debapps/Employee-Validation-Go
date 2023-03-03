package main

import (
	"empvalid/utilities/csvhandle"
	"empvalid/utilities/emputils"
	"fmt"
	"strings"
)

func main() {

	programName := "Handling Employee File"
	fmt.Print("\n")
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("%60s\n", programName)
	fmt.Println(strings.Repeat("-", 100))

	// Input and Output paths of the data files.
	inFilePath := "./data/input/Employee_Data.CSV"
	outFilePath := "./data/output/Employee_Data.ftp"

	// Read the input file and get the employee list.
	byteStram := csvhandle.GetFileStream(inFilePath)
	recs := csvhandle.GetCSVData(byteStram)
	empList := emputils.GetFileEmpList(recs)

	// Traverse Employee List. Checks for valid employee record.
	// Writes the valid records into output file.
	writer := csvhandle.GetCSVWriter(outFilePath)

	for _, emp := range empList {
		if emp.ValidateEmpFields() {
			fmt.Printf("\n\nWritting Valid Employee into File - %s \n", outFilePath)
			emp.ShowEmp()
			writer.WriteCSVData(emp.FormatEmpData())
		} else {
			fmt.Println("\nSkipping Invalid Employee -")
			emp.ShowEmp()
		}
	}

}
