package main

import "fmt"

func main() {
	// var ptr *int

	var ptr *int
	var a int = 10
	ptr = &a // referancing pointer to a

	fmt.Println("Value of a is :", a)
	fmt.Println("value of pointer is", ptr)
	fmt.Println("Value stored at pointer is :", *ptr) // dereferancing pointer

	ModifyValue(ptr)
	fmt.Println("Value of a after modifying through pointer is :", a)

}

func ModifyValue(ptr *int) {
	*ptr++
}
