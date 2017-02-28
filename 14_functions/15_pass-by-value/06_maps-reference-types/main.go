package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Some text"]) // 42
}

func changeMe(z map[string]int) {
	z["Some text"] = 42
}
