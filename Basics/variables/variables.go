package main

import "fmt"

// global variable
var globalVar string = "I am a global variable"

func main() {

	// local variable using :=
	localVar := "I am a local variable"

	fmt.Println(globalVar)
	fmt.Println(localVar)

	showScope()
}

func showScope() {
	// local to this function
	innerVar := "I am inside showScope function"

	fmt.Println(innerVar)
	fmt.Println(globalVar)
}
