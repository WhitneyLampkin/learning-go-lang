package variables

import "fmt"

func MostCommonVariableDeclaration() {
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)
}

func DeclaringVariables() {
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

func InitializingVariables() {
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
