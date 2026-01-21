package main

import "fmt"

// Struct declaration (global)
type User struct {
	Name string
	Age  int
}

// Method with value receiver
func (u User) Info() string {
	return fmt.Sprintf("%s (%d)", u.Name, u.Age)
}

// Method with pointer receiver
func (u *User) UpdateAge(age int) {
	u.Age = age
}

// Embedded struct
type Employee struct {
	User
	ID int
}

func main() {

	// Initialization
	u1 := User{Name: "Pratiksha", Age: 25}
	u2 := User{Name: "Pratiksha", Age: 25}

	// Access fields-----
	fmt.Println(u1.Name)

	// Compare structs
	fmt.Println("Equal:", u1 == u2)

	// Method calls
	fmt.Println(u1.Info())
	u1.UpdateAge(26)
	fmt.Println("Updated Age:", u1.Age)

	// Anonymous struct
	temp := struct {
		Task string
	}{
		Task: "Temporary Job",
	}
	fmt.Println(temp.Task)

	// Embedding
	emp := Employee{
		User: User{"Pratiksha", 26},
		ID:   101,
	}
	fmt.Println(emp.Name, emp.ID)
}
