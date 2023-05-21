package main

import "fmt"

func main() {
	//To simply remember what does the defer do we can remember "FIRST IN LAST OUT" but only with the  line with defer keyword
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
