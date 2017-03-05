package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	// student[0] = "new text"
	fmt.Println(student)
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println(students)
	fmt.Println(student == nil)
}
