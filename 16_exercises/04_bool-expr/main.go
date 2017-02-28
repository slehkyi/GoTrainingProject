package main

import "fmt"

func main() {
	v := (true && false) || (false && true) || !(false && false)
	fmt.Println(v)
}
