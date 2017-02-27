package main

import "fmt"

func main()  {

	myFriendsName := "Tim"

	switch {
	case len(myFriendsName) == 3:
		fmt.Println("Wassup my friend's name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Marcus", myFriendsName == "Jenny":
		fmt.Println("Wassup Marcus or Jenny")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}
