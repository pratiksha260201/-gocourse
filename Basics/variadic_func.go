package main

import "fmt"

func main() {

	//...  Ellipsis
	// fmt.Println("The sum is :", sum()) // if we pass npothing it will return 0
	// fmt.Println("The sum is :", sum(1, 2, 3, 4, 5))

	// statement, result := sum("The sum is:", 1, 2, 3, 4, 5)
	// fmt.Println(statement, result)
	// fmt.Println(sum("", 1, 2, 3, 4)) // if we dont want to print any statement use empty string
	// fmt.Println(sum("Total is:", 10, 20, 30))

	seq1, total := sum(1, 45, 34, 45, 67, 89)
	fmt.Println(seq1, total)

	seq2, total := sum(2, 75, 34, 45, 67, 89)
	fmt.Println(seq2, total)

	numbers := []int{10, 20, 30, 40}

	seq3, total := sum(3, numbers...) // unpacking the slice ... (Ellipsis) will helps to unpack the slice
	fmt.Println(seq3, total)

}

// func sum(nums ...int) int {
// 	total := 0
// 	for _, V := range nums {
// 		total += V

// 	}
// 	return total
// }

// func sum(returnStatement string, nums ...int) (string, int) {
// 	total := 0
// 	for _, V := range nums {
// 		total += V

// 	}
// 	return returnStatement, total
// }

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, V := range nums {
		total += V

	}
	return sequence, total
}
