package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")
	//There is no inheritance in golang

	aryan := User{"Aryan", "aryan@go.dev", true, 16}
	fmt.Println(aryan)
	fmt.Printf("aryan details are: %+v\n", aryan)
	fmt.Printf("Name is %v and Email is %v.\n", aryan.Name, aryan.Email)
	aryan.GetStatus()
	aryan.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is the user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
