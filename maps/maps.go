package maps

import "fmt"

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
