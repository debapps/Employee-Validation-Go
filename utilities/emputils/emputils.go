package emputils

import (
	"empvalid/models/employee"
)

func GetFileEmpList(empData [][]string) []employee.Employee {
	// Remove first element of employee data as it is header.
	empData = append(empData[:0], empData[1:]...)

	// Create an empty employee list.
	empList := []employee.Employee{}

	// Traverse the Employee data.
	for _, emp_data := range empData {
		// Create employee object.
		emp := employee.Employee{
			EmployeeID:   emp_data[0],
			FirstName:    emp_data[1],
			LastName:     emp_data[2],
			AddressLine1: emp_data[3],
			AddressLine2: emp_data[4],
			City:         emp_data[5],
			State:        emp_data[6],
			ZipCode:      emp_data[7]}

		// Append the Employee object to the list.
		empList = append(empList, emp)
	}

	return empList
}
