package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Method with pointer receiver
func (u *User) updateAge() {
	u.Age += 1
}

// Function using pointer
func updateValue(num *int) {
	*num = *num + 10
}

func main() {
	value := 20
	ptr := &value

	fmt.Println("Value:", value)
	fmt.Println("Value using pointer:", *ptr)

	*ptr = 30
	fmt.Println("Updated value:", value)

	updateValue(&value)
	fmt.Println("After function call:", value)

	user := User{Name: "abc", Age: 25}
	user.updateAge()
	fmt.Println("Updated user age:", user.Age)
}
