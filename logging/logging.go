package logging

import (
	"fmt"
	"log"
	"os"
	//"github.com/rs/zerolog"
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
	// Provides similar functionality as LogFatalMsg()
	// Panic provides the log message AND stacktrace
	log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}

func LogSetPrefix() {
	// log.SetPrefix is used to add a prefix to the log message
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}

func LogToFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")
}

func LogWithFramework() {
	/* zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Print("Hey! I'm a log message!")
	}

	func LogWithContext() {
		/* zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		zerolog.log.Debug().
			Int("EmployeeID", 1001).
			Msg("Getting employee information")

		zerolog.log.Debug().
			Str("Name", "John").
			Send()
	*/
}
