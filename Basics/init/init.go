package main

import "fmt"

//single unit func

// func init() {
// 	fmt.Println("Init function called")
// }

//multiple unit funcs

func init() {
	fmt.Println("FIRST Init function called")
}
func init() {
	fmt.Println("SECOND Init function called")
}

func init() {
	fmt.Println("THIRD Init function called")
}

//go automatically calls init before main
func main() {

	fmt.Println("Main function called")
}
