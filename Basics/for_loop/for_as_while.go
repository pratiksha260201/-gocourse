package Basics

import "fmt"

func main() {
	// i := 0
	// for i <= 5 {
	// 	println(i)
	// 	i++
	// }

	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println(sum)
	// 	if sum >= 30 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("odd number:", num)
		num++
	}

}
