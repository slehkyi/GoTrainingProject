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
				fmt.Println(a*b*c)
			}
		}
	}
}

/*
Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
 */
