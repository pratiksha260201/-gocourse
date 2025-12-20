package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	//Welcome message
	println("Welcome to the Guessing Game!")
	println("I have selected a number between 1 and 100.")
	println("Can you guess what it is?")

	var guess int

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number:", target)
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}

}
