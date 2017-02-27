package main

import "fmt"

func main() {
	switch "Jeny" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Mehdi":
		fmt.Println("Wassup Mehdi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
	no default fallthrough
	fallthrough is optional
	-- you can specify fallthrough by explicitly stating it
	-- break isn't needed like in other languages
*/
