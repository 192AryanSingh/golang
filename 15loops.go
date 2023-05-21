package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturaday"}

	fmt.Println(days)
	// using for loop method 1
	//	for d := 0; d< len(days); d++{
	//  fmt.Println(days[d])
	//}
	//method 2 for using loops
	//for i := range days {
	//fmt.Println(days[i])
	//}
	//method 3 for using loops initailizing indexes
	for index, day := range days {
		fmt.Printf("Index is %v and the value is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			//break
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
