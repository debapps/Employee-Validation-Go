package employee

import (
	"fmt"
	"strconv"
	"strings"
)

type Employee struct {
	EmployeeID, FirstName, LastName, AddressLine1, AddressLine2, City, State, ZipCode string
}

// This function shows employee details.
func (e *Employee) ShowEmp() {
	fmt.Printf("\nEmployee Details:\nEmployee ID - %s \nName - %s %s \nAddress Line 1 - %s \nAddress Line 2 - %s \nCity - %s \nState - %s \nZip Code - %s\n\n", e.EmployeeID, e.FirstName, e.LastName, e.AddressLine1, e.AddressLine2, e.City, e.State, e.ZipCode)
}

// This function validates Employee ID based on following criteria:
// 1. Should be 10 characters.
// 2. Should be digit only.
func (e *Employee) ValidateEmpId() bool {
	var validEmpID, isNumeric, is10Char bool

	if a, _ := strconv.Atoi(e.EmployeeID); a > 0 {
		isNumeric = true
	}

	if len(e.EmployeeID) <= 10 {
		is10Char = true
	}

	if isNumeric && is10Char {
		validEmpID = true
	}

	return validEmpID
}

// This function validates Zip code based on following criteria:
// 1. Should be 6 characters.
// 2. Should be digit only.
func (e *Employee) ValidateZip() bool {
	var validEmpID, isNumeric, is10Char bool

	if a, _ := strconv.Atoi(e.ZipCode); a > 0 {
		isNumeric = true
	}

	if len(e.ZipCode) == 6 {
		is10Char = true
	}

	if isNumeric && is10Char {
		validEmpID = true
	}

	return validEmpID
}

// This function validates all employee data fields.
func (e *Employee) ValidateEmpFields() bool {
	// Checks the fields contains any special characters or not.

	isValidFName := !checkSpclChar(e.FirstName)
	isValidLName := !checkSpclChar(e.LastName)
	isValidAddr1 := !checkSpclChar(e.AddressLine1)
	isValidAddr2 := !checkSpclChar(e.AddressLine2)
	isValidCity := !checkSpclChar(e.City)
	isValidState := !checkSpclChar(e.State)

	isValidEmp := e.ValidateEmpId() && e.ValidateZip() && isValidFName && isValidLName && isValidAddr1 && isValidAddr2 && isValidCity && isValidState

	return isValidEmp
}

// This function formats the employee data and returns the string slice of the employee details.
func (e *Employee) FormatEmpData() []string {
	// Format the employee data as follows:
	// Employee ID should be 10 character string padded with zero.
	// First Name should be 20 character string.
	// Last Name should be 30 character string.
	// Address Line 1 should be 50 character string.
	// Address Line 2 should be 50 character string.
	// City should be 20 character string.
	// State should be 20 character string.

	e.EmployeeID = fmt.Sprintf("%010s", e.EmployeeID)
	e.FirstName = fmt.Sprintf("%-20s", e.FirstName)
	e.LastName = fmt.Sprintf("%-30s", e.LastName)
	e.AddressLine1 = fmt.Sprintf("%-50s", e.AddressLine1)
	e.AddressLine2 = fmt.Sprintf("%-50s", e.AddressLine2)
	e.City = fmt.Sprintf("%-20s", e.City)
	e.State = fmt.Sprintf("%-20s", e.State)

	return []string{e.EmployeeID, e.FirstName, e.LastName, e.AddressLine1, e.AddressLine2, e.City, e.State, e.ZipCode}
}

// This function checks if there is a special character in the input string.
func checkSpclChar(inStr string) bool {
	// List of special characters.
	var specialChars = []rune{',', ';', ':', '"', '\'', '\\', '/', '{', '}', '|', '[', ']', '+',
		'=', '-', '_', '(', ')', '*', '&', 'ˆ', '%', '$', '#', '@', '!', '˜', '`'}

	var isContainSpecial bool
	for _, char := range specialChars {
		isContainSpecial = strings.ContainsRune(inStr, char)
		if isContainSpecial {
			break
		}
	}

	return isContainSpecial
}
