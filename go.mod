module main

go 1.19

replace empvalid/models/employee => ./models/employee

require (
	empvalid/models/employee v0.0.0-00010101000000-000000000000
	empvalid/utilities/csvhandle v0.0.0-00010101000000-000000000000
	empvalid/utilities/emputils v0.0.0-00010101000000-000000000000
)

replace empvalid/utilities/csvhandle => ./utilities/csvhandle

replace empvalid/utilities/emputils => ./utilities/emputils
