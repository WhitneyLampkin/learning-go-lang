# README.md

## About

Repository to store code snippets and notes from the following Microsoft Learn learning path: https://docs.microsoft.com/en-us/learn/paths/go-first-steps/.

## Running the App

1. cd helloworld
1. go build
1. go run main.go 1,2
    - Note: One of the functions expects 2 int args to be passed as seen by "1,2" above, but any int values can be used.

## Topics

1. Introduction
    - What is Go?
    - Exercise - Install GO
    - Exercise - Explore the GO Playground
    - Exercise - Install Visual Studio Code and the GO extension
    - Exercise - Hello World
    - Knowledge Check
    - Summary
1. Understand how to use packages, variables, and functions in GO
    - Introduction
    - Declare and use variables
    - Learn about basic data types
    - Create functions
    - Learn about packages
    - Knowledge check
    - Summary
1. Use control flows in GO
    - Introduction
    - Test conditions with if/else expressions
    - Control flow with switch statements
    - Loop through data with for expressions
    - Control with defer, panic, and recover functions
    - Exercise - Use control flows in GO
    - Solutions - Control flow exercises
    - Knowledge check
    - Summary
1. Use data types and structs, arrays, slices, and maps in GO
    - Introduction
    - Exercise - Use arrays
    - Exercise - Explore slices
    - Exercise - Use maps
    - Exercise - Use structs
    - Challenge - Data types
    - Solution - Data types challenge
    - Knowledge check
    - Summary
1. Implement error handling and logging in GO
    - Introduction
    - Learn how to handle errors in GO
    - Learn how to log in GO
    - Knowledge check
    - Summary
1. Use methods and interfaces in GO
    - Introduction
    - Use methods in GO
    - Use interfaces in GO
    - Challenge - Methods and interfaces
    - Solution - Methods and interfaces
    - Knowledge check
    - Summary
1. Learn how concurrency works in GO
    - Introduction
    - Learn about goroutines
    - Use channels as a communication mechanism
    - Learn about buffered channels
    - Challenge
    - Solution
    - Knowledge check
    - Summary
1. Write and test a program in GO
    - Introduction
    - Outline the online bank project
    - Get started with writing tests
    - Write the bank core package
    - Write hte bank API
    - Challenge - Complete the bank project functionality
    - Solution
    - Knowledge check
    - Summary

## Lessons learned from 'Take your first steps with Go'

1. Be mindful of where packages/modules are created in the directory
1. Be sure to open the scr folder from the correct $GOPATH path (This can break imports.)
1. Use 'go build' to automatically add all dependencies to go.mod
    - go get -u github.com/whitneylampkin/learning-go-lang/[FOLDER]@master
1. [ISSUE] go run main.go resulted in error: ".\main.go:128:2: undefined: errorhandling.HandleError"
    - [SOLUTION] go clean -modcache --> go build --> go run main.go
1. Adding new packages to an EXISTING project.
    - Create new folder
    - Add *.go file w/ "package [PACKAGE_NAME]" as the first line
    - cd into the package's directory
    - Create a module for the package with "go mod init [PACKAGE_NAME]
    - "go mod tidy"
    - Go to project's directory
        - go get [PACKAGE_NAME] to download the package
