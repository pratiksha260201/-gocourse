package main

import (
	"fmt"
	"sort"
)

func main() {

	// 1️⃣ Creating slice using slice literal
	nums := []int{10, 20, 30}
	fmt.Println("Initial slice:", nums)

	// 2️⃣ Creating slice using make
	emptySlice := make([]int, 0)
	fmt.Println("Empty slice:", emptySlice)

	// 3️⃣ Adding elements using append
	nums = append(nums, 40, 50)
	fmt.Println("After append:", nums)

	// 4️⃣ Modifying slice element
	nums[1] = 25
	fmt.Println("After modification:", nums)

	// 5️⃣ Creating slice from existing slice
	subSlice := nums[1:4]
	fmt.Println("Sub slice:", subSlice)

	// 6️⃣ Iterating over slice using range
	fmt.Println("Iterating over slice:")
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 7️⃣ Zero value of slice
	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice)
	fmt.Println("Is nil slice nil?", nilSlice == nil)

	// 8️⃣ Multidimensional slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multidimensional slice:", matrix)

	// 9️⃣ Sorting a slice
	sort.Ints(nums)
	fmt.Println("Sorted slice:", nums)
}
