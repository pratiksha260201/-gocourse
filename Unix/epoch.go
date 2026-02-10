package main

import (
	"fmt"
	"time"
)

func main(){

	t := time.Now()
	fmt.Println(t)

	UnixTime := t.Unix()
	fmt.Println("unix time",UnixTime)

	ConvertToIst := time.Unix(UnixTime,0)
	fmt.Println(ConvertToIst)

	//formtted IST 

	FormtString := ConvertToIst.Format("02 Jan 2006")
	fmt.Println("Formtted String:",FormtString)



}