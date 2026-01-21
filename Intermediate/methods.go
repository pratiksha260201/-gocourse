package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type shape struct {
	Rectangle
}

//method with value receiver

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

//method with pointer receiver
//if you want to modify the receiver instance and want to avoid copy large struct we will use pointer receiver

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

//method can be of any type

type Myint int

func (m Myint) isPositive() bool {
	return m > 0
}

//func or method without creating an instance we only create instance if we want to extract and modifying value, performing operation of that type
func (Myint) WelcomeMessage() string {
	return "Welcome to Myint type"
}

func main() {
	rect := Rectangle{length: 10, width: 5}
	area := rect.Area()
	fmt.Println("Area of the rectangle", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of the rectangle with factor of 2", area)

	num := Myint(-5)
	num1 := Myint(7)
	fmt.Println(num.isPositive())
	fmt.Println(num1.isPositive())

	fmt.Println(num1.WelcomeMessage())

	s := shape{Rectangle: Rectangle{length: 10, width: 6}}
	fmt.Println(s.Area())

}
