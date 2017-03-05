package main

import "fmt"

func main() {
	student := []string{"text"}
	students := [][]string{}
	student[0] = "new text"
	fmt.Println(student)
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println(students)
	fmt.Println(student == nil)
}
