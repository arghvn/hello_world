package main

//Sometimes we’ll want to build our programs into binaries. We can do this using go build.

//We can then execute the built binary directly.

import "fmt"

func main() {
	fmt.Println("hello world")
}

// if we type go run or go build hello_world in terminal ,out put :

//hello_world

// if we type ls , output :

//hello_world  hello_world.go

// if we type ./hello_world ,output :

//hello world

//more details :

//To start a program, a name must always be placed as the name of the package at the beginning of the file,
// such as the main package. You can use different names for Choose the name of your package,
//but the main file on which the executable command is issued must be the main package name.
//Note that when the program starts, only the code that we put in the main function will be executed,
// so it is better to use our code as Write the function outside it and call inside the main function.
