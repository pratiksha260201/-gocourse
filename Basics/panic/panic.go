package main

import (
	"fmt"
)

func main() {
	//panic(interface{})

	process(10)
	process(-5)

}

func process(i int) {
	defer fmt.Println("defferd statement 1")
	defer fmt.Println("defferd statement 2")
	if i < 0 {
		fmt.Println("before panic")
		panic("Negative value not allowed")

		//anything after panic will not be executed

		// fmt.Println("after panic")
		// defer fmt.Println("defferd statement 3")
	}

	fmt.Println("Proccesing with valid input", i)
}
