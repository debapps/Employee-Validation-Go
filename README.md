# Employee Data Validation

The purpose of the project is to validate employee CSV file data and generate the employee CSV file with valid data for FTP.

The input file of the project is to be taken from Amazon S3 bucket - **employee-data-raw**. The output file is to be stored in another S3 bucket - **employee-data-ftp**.

## Requirements

There are following requirements of the project.

1. Two Amazon S3 buckets needs to be created as follows:

   - **employee-data-raw**: Input Bucket.
   - **employee-data-ftp**: Output Bucket.

2. The input S3 bucket should have a CSV file called **Employee_Data.CSV** having following structure:

   - Employee ID
   - First Name
   - Last Name
   - Address Line 1
   - Address Line 2
   - City
   - State
   - Zip Code

3. The following validation needs to be performed on the CSV data of the input file:

   a. Employee ID should be 10 digit only. Sometimes padded with zero from the left.

   b. Zip Code should be 6 digit only.

   c. First Name, Last Name, Address Line 1, Address Line 2, City, State should NOT contain any special character.

4. If any of the above validation fails, the employee records should not be written into the output file.

5. Convert the each employee record into fixed width file having following column length:

   - Employee ID -> 10
   - First Name -> 20
   - Last Name -> 30
   - Address Line 1 -> 50
   - Address Line 2 -> 50
   - City -> 20
   - State -> 20
   - Zip Code -> 6

6. Upload the output file with name: **Employee-Data-Fixed-Width.txt** into output S3 bucket.

## Project Structure

In the root folder, the Golang script - _main.go_ is the entry point of the project.

The folder structures of the project is as follows:

### Data Model

---

- **models/employee** - Contains Employee Structure. The script file - _employee.go_. Module Name - _empvalid/models/employee_.

### Utility Function

---

- **utilities/emputils** - Contains Employee Utility function. The script file - _emputils.go_. Module Name - _empvalid/utilities/emputils_.

- **utilities/csvhandle** - Contains CSV File Handling Utility functions. The script file - _csvhandle.go_. Module Name - _empvalid/utilities/csvhandle_.

- **utilities/awsS3handle** - Contains Amazon S3 Bucket Handling Utility functions. The script file - _awsS3handle.go_. Module Name - _awsS3handle_.

### Input/Output Data

---

- **data/input** - Contains input data file for local testing.

- **data/output** - Contains the output file before uploading into Amazon S3 output bucket.

## Run Script

`go run main.go`

OR

`go run .`
