package main

import "fmt"

/*
	Variables: The following code will throw an error because variables are declared but not used.

	func main() {
		firstName, lastName := "John", "Doe"
		age := 32
	}

	./main.go:4:2: firstName declared but not used
	./main.go:4:13: lastName declared but not used
	./main.go:5:2: age declared but not used

	Data Types: Go is a strongly typed language so variables will only take values with the correct data type.

	- Basic types: numbers, strings, and booleans
	- Aggregate types: arrays and structs
	- Reference types: pointers, slices, maps, functions, and channels
	- Interface types: interface
*/

func main() {
	fmt.Println("Hello World!")

	mostCommonVariableDeclaration()
	declaringVariables()
	initializingVariables()
	integerDataType()
}

func mostCommonVariableDeclaration() {
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)
}

func declaringVariables() {
	// Declaring a single variable
	var myFirstName string
	myFirstName = "Whitney"
	fmt.Println(myFirstName)

	// Declaring multiple variables
	var hisFirstName, hisLastName string
	hisFirstName = "Regis"
	hisLastName = "Davis Jr."
	fmt.Println(hisFirstName, hisLastName)

	// Declaring multiple variable with different data types
	var herFirstName, herLastName string
	var herAge int
	herFirstName = "Kai"
	herLastName = "Lymore"
	herAge = 9
	fmt.Println(herFirstName, herLastName, herAge)

	// Simplified variable declaration
	var (
		firstName, lastName string
		age                 int
	)
	firstName = "First"
	lastName = "Last"
	age = 35
	fmt.Println(firstName, lastName, age)
}

func initializingVariables() {
	// Initializing variables with data types
	var (
		firstNameJohn string = "John"
		lastNameDoe   string = "Doe"
		age32         int    = 32
	)
	fmt.Println(firstNameJohn, lastNameDoe, age32)

	// Initializing variables without data types
	var (
		firstNameJohn2 = "John"
		lastNameDoe2   = "Doe"
		age33          = 33
	)
	println(firstNameJohn2, lastNameDoe2, age33)

	// Declaring and initializing variables in a single line
	var (
		firstNameJohn3, lastNameDoe3, age34 = "John", "Doe", 34
	)
	println(firstNameJohn3, lastNameDoe3, age34)
}

func integerDataType() {
	// Various integer types
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	/*
		- The following code throws an error because the two types don't match.
		- Must explicitly define the new type.

		var integer16 int16 = 127
			var integer32 int32 = 32767
			fmt.Println(integer16 + integer32)
		}

		Throws the following error: invalid operation: integer16 + integer32 (mismatched types int16 and int32)
	*/

	// Runes: "A rune is simply an alias for int32 data type. It's used to represent a Unicode character (or a Unicode code point)."
	rune := 'G'
	fmt.Println(rune)
}
