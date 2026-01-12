package main

import "fmt"

func main() {
	fmt.Println("hello")

	// func <name>(parameter list) returnType {
	// code block
	// return value
	// }

	// sum := add(6, 7)
	// fmt.Println("addition is:", add(5, 9))

	// greet := func() {
	// 	fmt.Println("This is anonymous function")
	// }
	// greet()

	// operation := add

	// result := operation(10, 15)
	// fmt.Println("result is:", result)

	//function as parameter
	res := applyOperation(3, 7, add)
	fmt.Println("result of applyOperation is:", res)

	//function as return type
	Add2 := getAddFunction(2)
	fmt.Println("addition of 2 and 4 is  :", Add2(4))

}

func add(a, b int) int {
	return a + b
}

//function that takes function as parameter

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)

}

//function that returns function

func getAddFunction(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
