package main

import "fmt"

func main()  {

	name := "Sergi"
	fmt.Println(name) // Sergio

	changeMe(name)

	fmt.Println(name)  // Sergio
}

func changeMe(z string) {
	fmt.Println(z)  // Sergio
	z = "Todd"
	fmt.Println(z)  // Todd
}