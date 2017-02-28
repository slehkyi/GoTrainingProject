package main

import "fmt"

func main() {
	func() {
		fmt.Println("Some text to the terminal")
	}()
}
