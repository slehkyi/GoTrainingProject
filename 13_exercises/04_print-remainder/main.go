package main

import "fmt"

func main() {

	var small int
	var larger int

	fmt.Print("Please, enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Now, enter a larger number: ")
	fmt.Scan(&larger)
	fmt.Println("Reminder after dividing will be: ", larger%small)
}
