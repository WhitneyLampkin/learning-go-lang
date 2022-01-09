package controlflows

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// If statements do not need ().
// Ternary if statements are not supported.
func IfStatements() {
	x := 27
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	if x%2 != 0 {
		fmt.Println(x, "is odd")
	}
}

/*
	This following code produces an error:

	if num := somenumber(); num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    fmt.Println(num)
*/

func CompoundIfStatements() {
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func givemeanumber() int {
	return -1
}

// If you run the following code several times, you'll see a different output every time. (But if you run the code in the Go Playground, you'll get the same result every time. That's one of the service's limitations.
func SwitchStatements() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	}

	fmt.Println("ok")
}

func SwitchMultipleExpressions() {
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
}

// This function uses a switch statement that returns multiple expressions based on the value matched.
func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

func SwitchInvokeFunction() {
	// The switch state can use a dynamic value from an invoked function.
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())
}

func SwitchRegEx() {
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}
}

// This function omits the conditions.
// Instead of passing a value to check against, the case is evaluated against a variable.
func OmitConditions() {
	rand.Seed(time.Now().Unix())
	r := rand.Float64()

	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}
}

// Using fallthrough to prevent breaks and run the code for the next case without evaluation.
// Be careful using this feature.
func SwitchBreaks() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
}
