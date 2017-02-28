package main

import "fmt"
import "math"

func main() {

	a := 0.0
	b := 0.0
	c := 0.0

	for a = 1; a < 500; a++ {
		for b = 1; b < 500; b++ {
			c = math.Sqrt(a*a+b*b)
			if a < b && b < c && a + b + c == 1000 {
				fmt.Println(a, b, c)
			}
		}
	}
}
