package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3}
	myNextSlice := []int{4, 5, 6}

	mySlice = append(mySlice, myNextSlice...)

	fmt.Println(mySlice)
}
