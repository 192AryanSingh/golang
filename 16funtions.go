package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions in golang")
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	proRes := proAdder(2, 5, 6, 7)
	fmt.Println("Pro result is: ", proRes)
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
func greeter() {
	fmt.Println("Namaste from golang")
}
func greeterTwo() {
	fmt.Println("Another method")
}
