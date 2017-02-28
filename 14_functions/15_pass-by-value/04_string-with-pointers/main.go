package main

import "fmt"

func main()  {

	name := "Sergio"
	fmt.Println(&name) // memory address - 0x82023c080

	changeMe(&name)

	fmt.Println(&name) // memory address still the same - 0x82023c080
	fmt.Println(name)  // Todd - value is different
}

func changeMe(z *string) {
	fmt.Println(z)  // memory address
	fmt.Println(*z) // Sergio
	*z = "Todd"
	fmt.Println(z)  // memory address is still the same
	fmt.Println(*z) // value is different - Todd, and value of the age is changed
}
