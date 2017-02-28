package main

import "fmt"

func twoReturns(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	num, even := twoReturns(7)
	fmt.Println(num, even)
}
