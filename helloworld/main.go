package main

import (
	"encoding/json"
	"fmt"
	/*

		Uncomment to use all of the functions in main()

		"os"

		"github.com/whitneylampkin/learning-go-lang/calculator"
		"github.com/whitneylampkin/learning-go-lang/controlflows"
		"github.com/whitneylampkin/learning-go-lang/datatypes"
		"github.com/whitneylampkin/learning-go-lang/forloops"
		"github.com/whitneylampkin/learning-go-lang/functions"
		"github.com/whitneylampkin/learning-go-lang/variables" */)

func main() {
	/* fmt.Println("Hello World!")

	// Referencing a local package
	fmt.Println("***Referencing a local package***")
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

	// Referencing 3rd-party packages
	fmt.Println("***Referencing 3rd party packages***")
	//fmt.Println(quote.Hello())

	// Variables, datatypes and functions
	fmt.Println("***Variables, datatypes and functions***")
	variables.MostCommonVariableDeclaration()
	variables.DeclaringVariables()
	variables.InitializingVariables()
	datatypes.IntegerDataType()
	datatypes.FloatDataType()
	datatypes.StringDataType()
	datatypes.DefaultValues()
	datatypes.TypeConversations()
	functions.MainFunctions()

	// Calling customSumFunction()
	fmt.Println("***Calling custom functions***")
	sum := functions.CustomSumFunction(os.Args[1], os.Args[2])
	println("Sum with a custom function:", sum)

	// Calling returnMultipleValues()
	fmt.Println("***Returning multiple values***")
	sum, mul := functions.ReturnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	// Discarding variables
	fmt.Println("***Discarding variables***")
	sum2, _ := functions.ReturnMultipleValues(os.Args[1], os.Args[2])
	fmt.Println("Sum2 (Discarding mul):", sum2)

	functions.PointerParametervalues()

	// Control Flow
	fmt.Println("***Control Flow***")
	controlflows.IfStatements()
	controlflows.CompoundIfStatements()
	controlflows.SwitchStatements()
	controlflows.SwitchMultipleExpressions()
	controlflows.SwitchInvokeFunction()
	controlflows.SwitchRegEx()
	controlflows.OmitConditions()
	controlflows.SwitchBreaks()
	controlflows.DeferExample()
	// The next line is commented out because it's stops execution.
	// controlflows.PanicHighLowExample()
	controlflows.RecoverExample()

	// Forloops
	forloops.BasicForLoop()
	forloops.FakeWhileLoop()
	forloops.InfiniteLoop()

	// Challenge: FizzBuzz
	Fizzbuzz()

	// Challenge: Find the primes (under 20)
	fmt.Println("Prime numbers less than 20:")

	for number := 2; number < 20; number++ {
		if findprimes(number) == true {
			fmt.Printf("%v ", number)
		}
	}

	// Challenge: Ask for a number, panic if negative
	val := 0
	fmt.Print("\r\nEnter number: ")
	fmt.Scanf("%d", &val)

	switch {
	case val < 0:
		panic("The program is crashing because you entered a negative number.")
	case val > 0:
		fmt.Println("You entered:", val)
	case val == 0:
		fmt.Println("0 is neither negative nor positive")

	}

	// Arrays
	DeclareArrays()
	InitializingArrays()
	EllipsisInArrays()
	ArrayLastPositionOnly()
	MultidimensionalArrays()
	DeclareAndInitializeASlice()
	SliceItems()
	ExtendedSlice()
	AppendSliceItems()
	RemoveSliceItems()
	CopySliceItems()
	CopySliceItemsFixed()


	DeclareAndInitializeMaps()
	AddMapItems()
	// Uncomment next line to see the error message
	// CreateFailedAddMapItemsCase()
	AccessMapItems()
	RemoveMapItems()
	LoopsThroughMaps()

	DeclareAndInitializeStructs()
	EmbeddingStructs()
	EncodeDecodeStructsJSON()

	*/

	CalculateFibonacci()
	//TranslateRomanNumerals()
}

func Fizzbuzz() {
	for i := 1; i <= 100; i++ {
		// TODO: Use switch case instead
		// Could also use i%15 for the first case
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
			fmt.Println(" - FizzBuzz")
		}
		if i%3 == 0 && i%5 != 0 {
			fmt.Println(i)
			fmt.Println(" - Fizz")
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println(i)
			fmt.Println(" - Buzz")
		}
	}

}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func DeclareArrays() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
}

func InitializingArrays() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}

func EllipsisInArrays() {
	cities2 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities 2:", cities2)
}

// Initializing an array in this way will create an array with a length of 100
// Only the last position will have a value set. The others will default to 0, which is the default for type int.
func ArrayLastPositionOnly() {
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}

func MultidimensionalArrays() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}

func DeclareAndInitializeASlice() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
}

func SliceItems() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
}

func ExtendedSlice() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter2 := months[3:6]
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
}

// How Slices differ from arrays:

// Items can be added to slices
func AppendSliceItems() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

// Must use s[i:p] to create a new slice with only the desired items.
func RemoveSliceItems() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}
}

// Changes to a slice changes the underlying array also as seen with CopySliceItems().
func CopySliceItems() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}

// Copy slices when you do not want to change the underlying array.
func CopySliceItemsFixed() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]

	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}

// Maps => hast tables (collection of key/value pairs)
// All keys and all values must share the same type but keys & values do not.

func DeclareAndInitializeMaps() {
	// Create empty map with make(): studentsAge := make(map[string]int)
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}

	fmt.Println(studentsAge)
}

func AddMapItems() {
	// Empty map
	studentsAge := make(map[string]int)

	// Adding items to the map
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println(studentsAge)
}

func CreateFailedAddMapItemsCase() {
	// Nil map
	var studentsAge map[string]int

	// Adding items fails
	// Panic msg: "panic: assignment to entry in nil map"
	// Therefore empty, not nil maps should always be used.
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println(studentsAge)
}

func AccessMapItems() {
	studentsAge := make(map[string]int)

	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println("Bob's age is", studentsAge["bob"])

	// Christy isn't in the map but it sends the default value instead of producing an error.
	// This is interesting...could it cause bugs?
	fmt.Println("Christy's age is", studentsAge["christy"])

	// Using a flag to see if Christy's age is in the map instead:
	age, exist := studentsAge["christy"]

	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
		fmt.Println("Is Christy's age in the map? ", exist)
	}
}

func RemoveMapItems() {
	// Create empty map
	studentsAge := make(map[string]int)

	// Add items to map
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	// Delete item
	delete(studentsAge, "john")

	// Christy isn't an item in the map but GO doesn't panic with the following:
	delete(studentsAge, "christy")

	fmt.Println(studentsAge)
}

func LoopsThroughMaps() {
	studentsAge := make(map[string]int)

	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// Use _ to ignore either the key or value of a map
	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}

	// Accessing only the items key
	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}

// Structs: collection of data fields in one structure
// Groups together related fields that could form a recod

// Declaring a struct
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func DeclareAndInitializeStructs() {
	// Declaring var with the struct as the type
	var john Employee
	fmt.Println(john)

	// Declare and initialize struct
	employee1 := Employee{1001, "John", "Doe", "Doe's Street"}
	employee1.ID = 1001
	fmt.Println(employee1.FirstName)

	employee2 := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee2.FirstName)

	// Use & operator to create a pointer to the struct
	// The struct becomes mutable
	employee1Copy := &employee1
	employee1Copy.FirstName = "David"
	fmt.Println("Original employee1: ", employee1)
	fmt.Println("Copy of employee1: ", employee1Copy)
}

func EmbeddingStructs() {
	type Person struct {
		ID        int
		FirstName string
		LastName  string
		Address   string
	}

	type Employee struct {
		Person
		ManagerID int
	}

	type Contractor struct {
		Person
		CompanyID int
	}

	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName)
}

func EncodeDecodeStructsJSON() {
	// The FirstName and Address fields will be renamed in the json result.
	type Person struct {
		ID        int
		FirstName string `json:"name"`
		LastName  string
		Address   string `json:"address,omitempty"`
	}

	type Employee struct {
		Person
		ManagerID int
	}

	type Contractor struct {
		Person
		CompanyID int
	}

	employees := []Employee{
		Employee{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		Employee{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	// json.Marshal() encodes a struct into JSON
	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	// json.Unmarshal() decodes a JSON string into a data structure
	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}

func CalculateFibonacci() []int {
	// Get users input
	userInput := 6

	if userInput < 2 {
		fmt.Println(make([]int, 0))
		return make([]int, 0)
	}

	fibonacciSequence := make([]int, userInput)

	fibonacciSequence[0], fibonacciSequence[1] = 1, 1

	for i := 2; i < userInput; i++ {
		fibonacciSequence[i] = fibonacciSequence[i-1] + fibonacciSequence[i-2]
	}

	fmt.Println(fibonacciSequence)

	return fibonacciSequence
}

/* func TranslateRomanNumerals() {
	romanNumbers := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// Get user input
	userInput := "MCLX"

	var sum int

	// Sum user input
	for i := 0; i < len(userInput); i++ {
		value, exist := romanNumbers[userInput[i]]

		if exist {
			sum += value
		} else {
			panic("Invalid roman numeral value entered.")
		}
	}

	fmt.Println("Roman Number Sum: ", sum)
} */
