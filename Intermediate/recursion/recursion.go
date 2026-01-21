package main

import "fmt"

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(3))

	fmt.Println(SumOfDigits(6))
	fmt.Println(SumOfDigits(67))
	fmt.Println(SumOfDigits(4567))

}

func factorial(n int) int {
	//base case factorial 0 =1
	if n == 0 {
		return 1
	}
	//recursive case factorial n = n*factorial(n-1)
	return n * factorial(n-1)
	//n *(n-1)*(n-2)*...*1
}

func SumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + SumOfDigits(n/10)
}
