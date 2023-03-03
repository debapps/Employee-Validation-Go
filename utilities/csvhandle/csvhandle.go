package csvhandle

import (
	"encoding/csv"
	"os"
	"strings"
)

type CSVWriter struct {
	writer *csv.Writer
}

// This function writes the slice of strings into CSV file.
func (w *CSVWriter) WriteCSVData(data []string) {
	err := w.writer.Write(data)
	checkError(err)
	w.writer.Flush()
}

// This function generates the CSV file writer object from the given file name and returns it.
func GetCSVWriter(filePath string) CSVWriter {
	// Open the file.
	outFile, err := os.Create(filePath)
	checkError(err)

	// Create the CSV file writer.
	writer := csv.NewWriter(outFile)

	// Generate the CSVWriter object and return it.
	writeObj := CSVWriter{writer: writer}
	return writeObj
}

// This function gets the byte stream from a file.
func GetFileStream(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	checkError(err)
	return data
}

// This function gets the CSV data in form of slice of string slice from the byte stream.
func GetCSVData(byteStram []byte) [][]string {
	// Creates the CSV reader object from the dara bytes.
	reader := csv.NewReader(strings.NewReader(string(byteStram)))

	// Reads all the records.
	records, err := reader.ReadAll()
	checkError(err)

	return records
}

// This function checks the error and raise if it exists.
func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
