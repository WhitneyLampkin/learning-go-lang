package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

// Concurrency: composition of independent activies (goroutines in Go); simultaneously but in an autonomous way
// Sharing data between processes in a common problem with concurrent programs.
// Go passes data between processes using channels.
// Only one activity (go routine) has access to data at a time...avoiding race conditions.

// Goroutines
// Goroutine: concurrent activity in a lightweight thread (not the traditional OS threads).
// Several Goroutines can have functions that are called at the same time.
// The first Goroutine is main().
// To create another Goroutine, use `go` before the function.

/* func GoroutineExample() {
	login()
    go launch()

	// Example with anonymous functions
	login()
    go func() {
        launch()
    }()
} */

func UseGoroutines() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	/* Leaving the code as is may take 2s or more

		for _, api := range apis {
		_, err := http.Get(api)
		if err != nil {
			fmt.Printf("ERROR: %s is down!\n", api)
			continue
		}

		fmt.Printf("SUCCESS: %s is up and running!\n", api)
	} */

	// Using the checkAPI Goroutine to check the status of the apis concurrently.
	// The `continue` keyword is no longer necessary when running concurrently.
	// This change cuts the time down to 500ms
	for _, api := range apis {
		go checkAPI(api)
	}

	// Using channels to improve the check api performance
	useChannel(apis)
	useBufferedChannel()

	//  The checks now run so fast that all of the output doesn't have time to print.
	// Uncommenting this sleep will show them again.
	// time.Sleep(3 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: %s is down!\n", api)
		// `continue` isn't necessary anymore; we can use `return` to
		return
	}

	fmt.Printf("SUCCESS: %s is up and running!\n", api)
}

// Channels

/*
	// Create channel
	ch := make(chan int)

	// Send/Receive via a channel
	ch <- x // sends (or write) x through channel ch
	x = <-ch // x receives (or reads) data sent to the channel ch
	<-ch // receives data, but the result is discarded

	// Close Channel
	close(ch)
*/

// Rewriting the checkAPI() function using channels
func checkAPIChannel(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		// Sprintf is used instead of Printf because we don't want to print, just send to
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	// (SEND/WRITE) <- ch <- (READ/RECEIVE)
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

// Send/Receive from channel
// Unbuffered Channels (Go's default)
func useChannel(apis []string) {
	ch := make(chan string)

	for _, api := range apis {
		go checkAPIChannel(api, ch)
	}

	fmt.Print(<-ch)
	// Program stopped listening at the last line.
	// Uncommenting the following line proves it:
	// fmt.Print(<-ch)
	// The above line will need to be repeated for all API checks to print to output.
	// However, this will still block the program from exiting.
	// Fix with;
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
}

// Buffered Channels
func send(ch chan string, message string) {
	ch <- message
}

func useBufferedChannel() {
	size := 4

	// size := 2 - causes a deadlock (causes to the send function are sequential)
	/*
		Fix with:

		size := 2
		ch := make(chan string, size)
		send(ch, "one")
		send(ch, "two")
		go send(ch, "three")
		go send(ch, "four")
		fmt.Println("All data sent to the channel ...")

		for i := 0; i < 4; i++ {
			fmt.Println(<-ch)
		}

		fmt.Println("Done!")
	*/

	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	send(ch, "three")
	send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}

/*
Channel Direction Example:

package main

import "fmt"

func send(ch chan<- string, message string) {
    fmt.Printf("Sending: %#v\n", message)
    ch <- message
}

func read(ch <-chan string) {
    fmt.Printf("Receiving: %#v\n", <-ch)
}

func main() {
    ch := make(chan string, 1)
    send(ch, "Hello World!")
    read(ch)
}

*/

/*
	Multiplexing:

	package main

	import (
		"fmt"
		"time"
	)

	func process(ch chan string) {
		time.Sleep(3 * time.Second)
		ch <- "Done processing!"
	}

	func replicate(ch chan string) {
		time.Sleep(1 * time.Second)
		ch <- "Done replicating!"
	}

	func main() {
		ch1 := make(chan string)
		ch2 := make(chan string)
		go process(ch1)
		go replicate(ch2)

		for i := 0; i < 2; i++ {
			select {
			case process := <-ch1:
				fmt.Println(process)
			case replicate := <-ch2:
				fmt.Println(replicate)
			}
		}
	}
*/
