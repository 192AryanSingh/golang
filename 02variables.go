package main

import "fmt"

//when the first letter of the funtion is capital then it's behaves public just like in c++
const LoginToken string = "rreggfgfg"

func main() {
	var username string = "Aryan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type : %T \n", isLoggedIn)
	var smallval int = 256
	fmt.Println(smallval)
	fmt.Printf("Variables is of type : %T \n", smallval)
	var smallfloat float64 = 256.322234342
	fmt.Println(smallfloat)
	fmt.Printf("Variables is of type : %T \n", smallfloat)
	//default values ans soome aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type : %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
	fmt.Printf("Variables is of type : %T \n", LoginToken)

}
