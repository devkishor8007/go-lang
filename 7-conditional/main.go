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
}