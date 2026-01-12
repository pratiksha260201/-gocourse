package main

import "fmt"

func main() {
	process()
	fmt.Println("Main function continues after recover")

}

func process() {
	defer func() {
		// if r := recover(); r != nil {

		//more readable way
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Process Start")
	panic("Someting wrong")
	fmt.Println("Process End")
}
