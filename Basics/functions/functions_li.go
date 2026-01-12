package main

import "fmt"

// Function declaration
// Takes two numbers and returns their sum
func add(a int, b int) int {
	return a + b
}

// Call by value example
func updateValue(x int) {
	x = x + 10
}

// Call by reference example
func updateReference(x *int) {
	*x = *x + 10
}

func main() {

	// Function calling with arguments
	result := add(5, 10)
	fmt.Println("Addition result:", result)

	// Call by value
	num := 20
	updateValue(num)
	fmt.Println("After call by value:", num)

	// Call by reference
	updateReference(&num)
	fmt.Println("After call by reference:", num)
}
