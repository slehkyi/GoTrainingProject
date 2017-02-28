package main

import "fmt"

func main()  {

	age := 44
	fmt.Println(&age) // memory address - 0x82023c080

	changeMe(&age)

	fmt.Println(&age) // memory address still the sane - 0x82023c080
	fmt.Println(age)  // 24 - value is different
}

func changeMe(z *int) {
	fmt.Println(z)  // memory address
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // memory address is still the same
	fmt.Println(*z) // value is different - 24, and value of the age is changed
}
