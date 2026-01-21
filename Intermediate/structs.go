package main

import (
	"fmt"
)

// always keep reminder that made struct and its method outside main function
type Person struct {
	firstName string
	LastName  string
	age       int
	address   Address // we can embed struct inside another struct
	PhoneCell         // Anonymous field
}

type Address struct {
	city    string
	country string
}

type PhoneCell struct {
	home string
	cell string
}

func main() {

	p1 := Person{
		firstName: "John",
		LastName:  "Doe",
		age:       30,

		//we can perform operation on struct als0
		address: Address{
			city:    "London",
			country: "U.K",
		},

		PhoneCell: PhoneCell{
			home: "5654457657656",
			cell: "7676786786786",
		},
	}

	//modity using another way

	// p1.address.city = "New York"

	//if we initialize filed and not assign value then it will take zero value for int and "" for string

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	fmt.Println("first Name:", p1.firstName)
	fmt.Println("first Name:", p2.firstName)
	fmt.Println("fullname:", p1.FullName())
	fmt.Println("age before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("after increment", p1.age)
	//accessing fields value from embed struct
	fmt.Println("City name is:", p1.address.city)

	//access complete struct
	fmt.Println("access complete struct", p1.address)
	//Anonymous Struct this are struct without name

	fmt.Println("acess home number from Phonecell", p1.cell) //PhoneCell is Anonymous filed hence directly but we cant acess city dircetly like this

	//comparing tow instance use == operator

	fmt.Println("p1 and p2 are equal ?:", p1 == p2)
	fmt.Println("p2 and p3 are equal ?:", p2 == p3)

	//user is anonymous struct

	user := struct {
		username string
		email    string
	}{
		username: "johndoe",
		email:    "john.doe@example.com",
	}

	fmt.Println("Username:", user.username)
	fmt.Println("Email:", user.email)

}

// method with struct value receiver
func (p Person) FullName() string {
	return p.firstName + " " + p.LastName
}

// to modity the fields within the method we use pointer receiver
func (p *Person) incrementAgeByOne() {
	p.age++
}
