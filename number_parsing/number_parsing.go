package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "12345"
	strcov, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	fmt.Println("Parsed integer", strcov)
	fmt.Println("Add 1 into int", strcov+1)

	//conver into specific size

	Strconint32, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("Error while Parsing into int32")
	}
	fmt.Println(Strconint32)

	//float conversation
	str1 := "4.3535"
	ConvToFloat, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("Float Convert value is %.2f\n", ConvToFloat)
	//.2 only return 2 values after decimal

	//Parse binary to decimal

	binary := "1010"
	//conver into base 10
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("decimal number is ", decimal)
}
