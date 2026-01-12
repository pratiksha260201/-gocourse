package main

import "fmt"

func main() {

	// addition1 := add(4, 6)
	// defer fmt.Println("The addition is:", addition1)

	// addition2 := add(10, 20)
	// fmt.Println("The addition is:", addition2)

	// process(10)

	//anonymous function with defer

	// defer func(a, b int) {
	// 	result := a + b
	// 	fmt.Println("The addition is:", result)
	// }(4, 6)

	fmt.Println("The addition is:", add(4, 6))

}

// func add(a, b int) int {
// 	return a + b
// }

//func with multi defer statements follow LIFO (Last In First Out) order
//if we pass arguments to defer statement those will be evaluated immediately

// func process(i int) {
// 	defer fmt.Println("Defer statement with argument:", i)
// 	defer fmt.Println("First defer statement")
// 	defer fmt.Println("Second defer statement")
// 	defer fmt.Println("Third defer statement")
// 	i++
// 	fmt.Println("Normal Statement")
// 	fmt.Println("after increment value of i is :", i)
// }

//defer can modify the return value only when function has named return variable

func add(a, b int) (result int) {
	defer func() {
		result += 10
	}()
	return a + b
}
