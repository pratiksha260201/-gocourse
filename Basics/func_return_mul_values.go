package main

import (
	"errors"
	"fmt"
)

func main() {

	q, r := divide(21, 7)
	fmt.Printf("quotient is: %d,reminder is: %d\n", q, r)

	// we can also use %V for geneal placeholder
	// fmt.Printf("quotient is: %v, reminder is: %v", q, r)

	res, err := compare(5, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}

}

func divide(a, b int) (int, int) {
	quotient := a / b
	reminder := a % b
	return quotient, reminder

}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "a is less than b", nil
	} else {
		return "", errors.New("a is equal to b")
	}

}
