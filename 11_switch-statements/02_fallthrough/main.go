package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Mehdi":
		fmt.Println("Wassup Mehdi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
		fallthrough
	default:
		fmt.Println("Have you no friends?")
	}
}
