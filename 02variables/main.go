package main

import "fmt"

const LoggedToken string = "gfhfjfj"     //public

func main() {
	//string
	var username string = "chandan"
	fmt.Println(username)
	fmt.Printf("username is a type of : %T \n",username )

    //boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("username is a type of : %T \n",isLoggedIn )

    //uint
	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("username is a type of : %T \n",smallVal )

	//float value
	var smallFloat float64 = 255.62346743732473287872
	fmt.Println(smallFloat)
	fmt.Printf("username is a type of : %T \n",smallFloat )

	//default values and aliases
	var anotherVariables int 
	fmt.Println(anotherVariables)
	fmt.Printf("username is a type of : %T \n",anotherVariables )

	// no var style 
	numberofPeople := 20000
	fmt.Println(numberofPeople)

	fmt.Println(LoggedToken)
	fmt.Printf("username is a type of : %T \n",LoggedToken)

}