package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to video on slices")

	var fruitlist = []string{"Apple", "Watermelon", "Grapes"}
	fmt.Printf("Type of fruitlist is : %T \n", fruitlist)

	fruitlist = append(fruitlist, "Banana", "Litchi")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[2:4])
	fmt.Println(fruitlist)


	highScores := make([]int , 4)

	highScores[0] = 123
	highScores[1] = 456
	highScores[2] = 789
	highScores[3] = 546

	highScores = append(highScores, 222, 333)

	fmt.Println(highScores)
	sort.Ints(highScores)    // for sorting integers
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))    // for boolean values
	fmt.Println(highScores)

	// remove a value from slices based on index
	var courses = []string{"react","native","python","golang","java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
