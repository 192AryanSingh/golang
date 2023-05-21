package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")
	//There is no inheritance in golang

	aryan := User{"Aryan", "aryan@go.dev", true, 16}
	fmt.Println(aryan)
	fmt.Printf("aryan details are: %+v\n", aryan)
	fmt.Printf("Name is %v and Email is %v", aryan.Name, aryan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
