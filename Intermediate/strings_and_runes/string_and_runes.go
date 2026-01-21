package main

import "fmt"

func main() {
	a := "Go"
	b := "Golang"
	text := "Go😊"

	// String comparison
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a < b:", a < b)

	// String vs Rune behavior
	fmt.Println("Byte length:", len(text))
	fmt.Println("Character length:", len([]rune(text)))

	for i, r := range text {
		fmt.Printf("Index %d: %c\n", i, r)
	}
}
