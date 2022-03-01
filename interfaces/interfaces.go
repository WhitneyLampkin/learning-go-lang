package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

// Interface: An abstract type that only includes methods a concrete type must possess or implement.
// Similar to a blueprint
// Interfaces make Go code more flexible and extensible.

func DeclearingInterfaces() {
	/*
		var s Shape = Square{3}
		fmt.Printf("%T\n", s)
		fmt.Println("Area: ", s.Area())
		fmt.Println("Perimeter:", s.Perimeter())
	*/

	var s Shape = Square{3}
	printInformation(s)

	c := Circle{6}
	printInformation(c)

	// Running the Stringer example
	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)
}

type Shape interface {
	// Any type that we want to consider a shape but have a Perimeter() and Area() function.
	Perimeter() float64
	Area() float64
}

// Interfaces don't use keywords for implementation.
// Simply satisfying the requires creates the concrete type from that interface.
type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

// Another concrete type, Circle, from interface.
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

// Stringer Example
type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

// Extending existing implementations
type customWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func writerExample() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Copy uses a Writer interface. This can be replaced with our own custom Write function.
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}
