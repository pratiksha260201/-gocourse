package main

import "fmt"

func main() {

	// 1️⃣ Map initialization using literal
	studentMarks := map[string]int{
		"Math":    85,
		"Science": 90,
	}

	// 2️⃣ Map initialization using make
	countryCode := make(map[string]string)
	countryCode["IN"] = "India"
	countryCode["US"] = "United States"

	// 3️⃣ Adding & modifying values
	studentMarks["English"] = 88
	studentMarks["Math"] = 92

	// 4️⃣ Retrieving value
	fmt.Println("Math marks:", studentMarks["Math"])

	// 5️⃣ Checking key existence
	value, exists := studentMarks["History"]
	if exists {
		fmt.Println("History marks:", value)
	} else {
		fmt.Println("History subject not found")
	}

	// 6️⃣ Iterating over map
	fmt.Println("Student Marks:")
	for subject, marks := range studentMarks {
		fmt.Println(subject, ":", marks)
	}

	// 7️⃣ Deleting key-value pair
	delete(studentMarks, "Science")
	fmt.Println("After deletion:", studentMarks)
}
