package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our shop")
	fmt.Println("Please give us rating between 1 to 4")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for giving us rating ," , input)

	newRating , err := strconv.ParseFloat(strings.TrimSpace(input),64)
	
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println("Thanks for adding 1 in our review ," , newRating+1)
	}
	

}