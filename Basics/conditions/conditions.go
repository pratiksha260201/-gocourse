package main

import "fmt"

func main() {
	age := 22
	score := 85
	day := "Sunday"

	if age >= 18 {
		fmt.Println("User is an adult")
	}

	if score >= 50 {
		fmt.Println("Exam passed")
	} else {
		fmt.Println("Exam failed")
	}

	if score >= 60 {
		if score >= 80 {
			fmt.Println("Grade: A")
		} else {
			fmt.Println("Grade: B")
		}
	}

	if age >= 18 && score >= 50 {
		fmt.Println("Eligible for next level")
	}

	if day == "Saturday" || day == "Sunday" {
		fmt.Println("It is a weekend")
	}
}
