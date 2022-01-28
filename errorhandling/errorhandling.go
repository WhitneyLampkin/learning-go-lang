package errorhandling

/*

Always check for errors, even if you don't expect them. Then handle them properly to avoid exposing unnecessary information to end users.
Include a prefix in an error message so you know the origin of the error. For example, you could include the name of the package and function.
Create reusable error variables as much as you can.
Understand the difference between using returning errors and panicking. Panic when there's nothing else you can do. For example, if a dependency isn't ready, it doesn't make sense for the program to work (unless you want to run a default behavior).
Log errors with as many details as possible (we'll cover how in the next section) and print out errors that an end user can understand.

*/

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func HandleError() {
	employee, err := getInformationReusableErr(1001)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	if err != nil {
		// return nil, err // Simply return the error to the caller.
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, nil
}

func getInformationRetry(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(1000)
		if err == nil {
			return employee, nil
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

var ErrNotFound = errors.New("Employee not found!")

func getInformationReusableErr(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
