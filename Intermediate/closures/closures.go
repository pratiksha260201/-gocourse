// package main

// import "fmt"

// func main() {

// 	// seq := adder()

// 	// fmt.Println("after adding 1 to i ", seq()) //1
// 	// fmt.Println("after adding 1 to i ", seq()) //2
// 	// fmt.Println("after adding 1 to i ", seq()) //3

// 	// seq2 := adder()

// 	// fmt.Println("after adding 1 to i in seq2 ", seq2()) //1
// 	// fmt.Println("after adding 1 to i in seq2 ", seq2()) //2

// 	substractor := func() func(int) int {
// 		coutdown := 99
// 		return func(x int) int {
// 			coutdown -= x
// 			return coutdown

// 		}

// 	}()
// 	fmt.Println("after substracting 1 ", substractor(1))       //98
// 	fmt.Println("after substracting again1  ", substractor(1)) //

// }

// // func adder() func() int {
// // 	i := 0
// // 	fmt.Println("Intitial value of i is :", i)
// // 	return func() int {
// // 		i++
// // 		fmt.Println("added 1 to i")
// // 		return i
// // 	}
// // }

package main

import "fmt"

func main() {

	// var i int = 0
	// i := 0
	// var i int
	// i = 0

	// counter := func() int {
	// 	i++
	// 	return i
	// }

	// fmt.Println(counter())

	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())

	counter1 := newCounter()
	fmt.Println(counter1())

}
func newCounter() func() int {
	var i int = 0
	return func() int {
		i++
		return i

	}
}
