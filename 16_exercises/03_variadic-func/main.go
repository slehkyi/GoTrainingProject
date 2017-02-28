package main

import "fmt"

func main() {
	x := largest(-1, -2, -13, -4, -25, -6, -7)
	fmt.Println(x)
}

func largest(x ...int) int {
	var lnum int
	for i, v := range x {
		if lnum < v || i == 0 {
			lnum = v
		}
	}
	return lnum
}
