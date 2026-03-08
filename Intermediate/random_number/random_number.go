package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {

	// fmt.Println(rand.Intn(10))
	// fmt.Println(rand.Intn(100) + 1)

	// fmt.Println(rand.Intn(5) + 6)

	// To generate random number between min and max:
	// rand.Intn(max-min+1) + min
	// between 10 and 20
	// rand.Intn(11) + 10

	// set seed
	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(val.Intn(5) + 1)

	fmt.Println(rand.Float64()) // generate random number between 0.0 to 1.0
	for {
		//show menu
		fmt.Println("Welcome to dice game")
		fmt.Println("1. roll the dice:")
		fmt.Println("2.Exit")
		fmt.Print("Enter the choice 1/2: \n")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("incorrect choice")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing, Good Byeee")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("you enter a %d and a %d \n", die1, die2)
		fmt.Println("total :", die1+die2)

		var PlayAgain string
		fmt.Println("if you want to play again y/n :")
		_, err = fmt.Scan(&PlayAgain)

		if err != nil || (PlayAgain != "y" && PlayAgain != "n") {
			fmt.Println("invalid assu,ing no")
		}

		if PlayAgain == "n" {
			fmt.Println("Thank for playing, byee")
		}
	}

}
