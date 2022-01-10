package forloops

import (
	"fmt"
	"math/rand"
	"time"
)

// Basic for loop syntax
func BasicForLoop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum of 1..100 is", sum)
}

// Use for loop without pre- and post-statements to create while loops
func FakeWhileLoop() {
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}

func InfiniteLoop() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writing inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

// Used for validation or to wait for external resources to become available
func ContinueStatements() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}
