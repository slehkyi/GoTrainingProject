package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}

	fmt.Println(myGreeting)
}
