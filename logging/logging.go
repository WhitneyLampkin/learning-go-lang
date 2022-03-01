package logging

import (
	"fmt"
	"log"
)

func LogStandardMsg() {
	// We could also log with fmt.Print() but log.Print() provides more functionality.
	log.Print("\r\nHey, I'm a log!")
}

func LogFatalMsg() {
	// Fatal logs - logs an error and then ends the program similarly to os.Exit(1)
	// log.Fatal("\r\nHey, I'm an error log!")
	// The next line won't run when the above line is uncommented because the program will exit.
	fmt.Print("\r\nCan you see me?")
}

func LogPanicMsg() {

}

func LogSetPrefix() {

}

func LogToFile() {

}

func LogWithFramework() {

}

func LogWithContext() {

}
