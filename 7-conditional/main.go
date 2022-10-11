package main

import ("fmt")

func main() {
	age := 18

	if age > 18 {
		fmt.Println("You can join on party")
	} else if age == 18 {
		fmt.Println("luckie you are in")
	} else {
		fmt.Println("age must be 18 or above 18")
	}

	// switch case

	switch age {
	case 15: 
	    fmt.Println("Not qualified")
	case 18:
		fmt.Println("you are in")
    case 20:
		fmt.Println("you are under qualified")
	}
}