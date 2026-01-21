package main

import (
	"fmt"
)

// Multiple return values with error handling
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Variadic function with simple parameter
func calculateTotal(name string, marks ...int) int {
	total := 0
	for _, mark := range marks {
		total += mark
	}
	fmt.Println("Student:", name)
	return total
}

func main() {

	// Error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	// Variadic function
	totalMarks := calculateTotal("Pratiksha", 80, 85, 90)
	fmt.Println("Total Marks:", totalMarks)

	// Anonymous function
	message := func(name string) {
		fmt.Println("Welcome,", name)
	}
	message("Golang")
}
