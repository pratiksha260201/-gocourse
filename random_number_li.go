package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("0-9:", rand.Intn(10))
	fmt.Println("1-10:", rand.Intn(10)+1)
	fmt.Println("5-15:", rand.Intn(11)+5)
	fmt.Println("Float:", rand.Float64())
}
