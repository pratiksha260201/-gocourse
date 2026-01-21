package main

import "fmt"

func groceryStore() {

	// Recover: Manager handling the issue
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Manager handled the issue:", err)
		}
	}()

	fmt.Println("Took shopping basket")

	// Defer: Basket must be returned
	defer fmt.Println("Returned shopping basket")

	// Panic: Unexpected system failure
	panic("Billing system crashed")

	fmt.Println("Billing completed")
}

func main() {
	groceryStore()
	fmt.Println("Store is still operating")
}
