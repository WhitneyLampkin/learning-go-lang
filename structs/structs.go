package structs

import (
	"encoding/json"
	"fmt"
)

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
