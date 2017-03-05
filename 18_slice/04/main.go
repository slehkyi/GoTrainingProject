package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Buenos dias!",
		"Bongiorno!",
		"Ohayo!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[3:] ")
	fmt.Println(greeting[3:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}
