package main

import "fmt"

func main() {
	marks := [5]int{78, 85, 90, 66, 88}

	fmt.Println("Accessing index and value:")
	for index, mark := range marks {
		fmt.Println("Index:", index, "Value:", mark)
	}

	fmt.Println("\nUsing blank identifier to ignore index:")
	for _, mark := range marks {
		fmt.Println("Value:", mark)
	}
}
