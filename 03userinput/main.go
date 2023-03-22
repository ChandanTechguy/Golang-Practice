package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome Bro"
	fmt.Println(welcome)

	 reader := bufio.NewReader(os.Stdin)
	 fmt.Println("Enter the rating for our Pizza:")

	 //comma ok // comma err

	 input,_ := reader.ReadString('\n')
	 fmt.Println("Thanks for rating = ",input)
	 fmt.Printf("Type for this rating is : %T",input)

}