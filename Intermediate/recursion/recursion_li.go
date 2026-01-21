package main

import "fmt"

// Direct Recursion
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Head Recursion
func printNumbersHead(n int) {
	if n == 0 {
		return
	}
	printNumbersHead(n - 1)
	fmt.Println(n)
}

// Tail Recursion
func printNumbersTail(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNumbersTail(n - 1)
}

// Indirect Recursion
func even(n int) {
	if n == 0 {
		fmt.Println("Even")
		return
	}
	odd(n - 1)
}

func odd(n int) {
	if n == 0 {
		fmt.Println("Odd")
		return
	}
	even(n - 1)
}

func intermediate() {
	fmt.Println("Factorial:", factorial(5))

	fmt.Println("\nHead Recursion:")
	printNumbersHead(3)

	fmt.Println("\nTail Recursion:")
	printNumbersTail(3)

	fmt.Println("\nIndirect Recursion:")
	even(4)
}
