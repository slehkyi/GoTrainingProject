package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	// fmt.Printf("%T \n", rune(foo))
}
