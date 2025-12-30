package main

import "fmt"

func main() {
	//Declare slice
	// 	//var sliceName[]dataType

	// var numbers []int
	// var numbers1 = []int{1,2,3,4,5}

	// numbers2 := []int{6,7,8,9,10}

	// //slice using make function
	// slice := make([]int, 5) //length 5

	a := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := a[1:4] //slicing array from index 1 to 3

	fmt.Println(sliceFromArray)

	slice1 := append(sliceFromArray, 7, 8)
	fmt.Println(slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println(sliceCopy)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

}
