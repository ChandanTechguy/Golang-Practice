package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[2] = "Litchi"
	fruitList[3] = "Banana"

	fmt.Println("Fruitlist is :", fruitList)
	fmt.Println("Fruitlist is :", len(fruitList))     // for arrays length

	var vegList = [3]string{"tomato","mushroom","potato"}
	fmt.Println("Veglist is :",len(vegList))

}