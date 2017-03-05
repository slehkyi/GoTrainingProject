package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	// var myGreeting = map[string]string{} -- works as well
	myGreeting["Tim"] = "Good morning!"
	myGreeting["Jenny"] = "Bonjour!"

	fmt.Println(myGreeting)
}
