package main

import "fmt"

func main() {
	customerNumber := make([]int, 3) // 3 is length & capacity
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 13

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning"
	greeting[1] = "Chiao"
	greeting[2] = "Hola"

	fmt.Println(greeting[2])
}
