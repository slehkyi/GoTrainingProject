package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string) {
	z[0] = "Todd"
	// z[1] = "Sergi"
	fmt.Println(z) // [Todd Sergi]
}
