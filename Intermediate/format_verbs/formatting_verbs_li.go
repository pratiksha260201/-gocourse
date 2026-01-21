package main

import "fmt"

type User struct {
	Name string
	Age  int
	Role string
}

func main() {
	// Basic values
	name := "Golang"
	version := 1.22
	active := true
	score := 89.456

	// Struct value
	user := User{
		Name: "Pratiksha",
		Age:  25,
		Role: "Backend Developer",
	}

	// Formatting basic values
	fmt.Printf("Language: %s\n", name)
	fmt.Printf("Version: %.2f\n", version)
	fmt.Printf("Active: %t\n", active)
	fmt.Printf("Score: %.1f\n", score)
	fmt.Printf("Type of version: %T\n\n", version)

	// Struct formatting
	fmt.Printf("User (default): %v\n", user)
	fmt.Printf("User (with fields): %+v\n\n", user)

	// Width and alignment
	fmt.Printf("|%10s|\n", name)  // Right aligned
	fmt.Printf("|%-10s|\n", name) // Left aligned

	// Using Sprintf
	msg := fmt.Sprintf("User %s logged in", user.Name)
	fmt.Println("\nMessage:", msg)
}
