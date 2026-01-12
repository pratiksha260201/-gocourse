package main

import (
	"fmt"
	"os"
)

func main() {

	//when exit is called defer statements will not be executed

	defer fmt.Println("defer 1")

	fmt.Println("Starting of the function ")

	//immediately terminates the program without running ant cleanup operations

	os.Exit(1)

	fmt.Println("This line will not be executed")

}
