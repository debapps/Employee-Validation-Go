package main

import (
	"awsS3handle"
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
	// inFilePath := "./data/input/Employee_Data.CSV"
	outFilePath := "./data/output/Employee_Data.ftp"

	// Input and Output bucket and object.
	inBucketName := "employee-data-raw"
	outBucketName := "employee-data-ftp"
	inObjectKey := "Employee_Data.CSV"
	outObjectKey := "Employee-Data-Fixed-Width.txt"

	// Read the input file from the S3 bucket object specified.
	s3Client := awsS3handle.InitS3Client()
	byteStream := s3Client.GetObjectContent(inBucketName, inObjectKey)

	// Read the input file and get the employee list.
	// byteStream := csvhandle.GetFileStream(inFilePath)
	recs := csvhandle.GetCSVData(byteStream)
	empList := emputils.GetFileEmpList(recs)
	// fmt.Println(empList)

	// Traverse Employee List. Checks for valid employee record.
	// Writes the valid records into output file.

	writer := csvhandle.GetCSVWriter(outFilePath)

	for _, emp := range empList {
		if emp.ValidateEmpFields() {
			fmt.Println("\n\nWritting Valid Employee")
			emp.ShowEmp()
			writer.WriteCSVData(emp.FormatEmpData())
		} else {
			fmt.Println("\nSkipping Invalid Employee -")
			emp.ShowEmp()
		}
	}

	// Upload the output file into S3 bucket.
	s3Client.UploadObject(outBucketName, outObjectKey, outFilePath)

	// Remove outfile.
	// fmt.Printf("\nRemoving %s ..", outFilePath)
	// err := os.Remove(outFilePath)
	// if err != nil {
	// 	panic(err.Error())
	// }

	fmt.Println("Completed!!")
}
