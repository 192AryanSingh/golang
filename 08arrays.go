package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))
	var vegyList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegy list is : ", vegyList)
	fmt.Println("vegy list is : ", len(vegyList))
}
