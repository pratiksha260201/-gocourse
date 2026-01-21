package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	counter1 := counter()
	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2

	counter2 := counter()
	fmt.Println(counter2()) // 1 (starts again from 0)
}
