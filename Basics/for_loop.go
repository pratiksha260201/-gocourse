package Basics

import "fmt"

func main() {

	//simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//ierating over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//break and continue

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break
		}
	}

	//Printing stars in a pattern
	rows := 5
	// Outer loop for each row
	for i := 1; i <= rows; i++ {
		//print spaces in inner loop
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := range 10 {
		i++
		fmt.Println(i)
	}

}
