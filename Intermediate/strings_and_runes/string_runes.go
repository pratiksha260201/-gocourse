package main

import (
	"fmt"
)

func main() {

	message := "Hello, Go!"
	message1 := "Hello,\nGo!"
	message2 := "Hello,\tGo!"
	message3 := "Hello,\rGo!"
	rawMessage := `Hello,\nGo!`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	//calculating length of string
	fmt.Println("Length of message is :", len(message))
	fmt.Println("Length of rawMessage is :", len(rawMessage))

	//concatenation of strings

	greeting := "Hello,"
	name := "Alice"
	fullGreeting := greeting + " " + name + "!"
	fmt.Println(fullGreeting)

	//string comparison
	str1 := "apple" // str1 is greater than str4 because lowercase letters have higher ASCII values than uppercase letters
	str2 := "banana"
	str3 := "app"
	str4 := "APPLE"

	fmt.Println("Is str1 equal to str2?", str1 == str2)
	fmt.Println("Is str1 less than str4?", str1 < str4)
	fmt.Println("Is str1 greater than str3?", str1 > str3)

	//iterating over string runes
	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		//hexadecimal value
		fmt.Printf("Hexadecimal value of character %c is %x\n", char, char)
	}

	//strings are immutable
	greetingWithName := greeting + name
	fmt.Println("Greeting with name:", greetingWithName)

	//there is no direct way to modify a string in Go

	//there is no character declaration in Go
	// var ch byte = 'A' // This is valid, but 'ch' is of type byte, not rune

	var ch rune = 'a' // 'ch' is of type rune (alias for int32)
	jch := '字'        // 'jch' is of type rune representing a Chinese character
	fmt.Println(ch)
	fmt.Println(jch)

	//actual value of ch and jch
	fmt.Printf("Character ch is: %c\n", ch)
	fmt.Printf("Character jch is: %c\n", jch)

	Converted_Str := string(ch)
	fmt.Println("Converted_Str is :", Converted_Str)
	//check type of Converted_Str
	fmt.Printf("Type of ch is : %T\n", ch)
	fmt.Printf("Type of Converted_Str is : %T\n", Converted_Str)

	const NIHANGO = "日本語" // Janpnese text
	jhello := "こんにちは"

	fmt.Println(NIHANGO)
	fmt.Println(jhello)

	//iterating over Japanese string runes
	for _, runeValue := range jhello {
		// fmt.Printf("%c\n", runeValue)
		fmt.Printf("%v\n", runeValue) // %v prints the default value of the variable

	}

	//we can print imogis also
	smile := '😊'
	fmt.Printf("%c\n", smile)
	fmt.Printf("%v\n", smile)

}
