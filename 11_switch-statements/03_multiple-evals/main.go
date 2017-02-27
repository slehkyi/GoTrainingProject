package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Some text for these names")
	case "Julian", "Susan":
		fmt.Println("Wassup guys")
	}
}
