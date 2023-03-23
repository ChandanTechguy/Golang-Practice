package main

import "fmt"

func main() {
	fmt.Println("Class on pointers")

	// var ptr *int   // * represents as pointer
	// fmt.Println("The value of pointer is ", ptr)

	myNumber := 20

	var ptr = &myNumber // reference = &
	fmt.Println("value of actual pointer is ",ptr)
	fmt.Println("value of actual pointer is ",*ptr)

	*ptr = *ptr +2
	fmt.Println("The value of pointer is ", myNumber)
}