package main

import ("fmt")

func main() {
	fmt.Println("Please enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("hii, %s! I am going!", name)
}

// to run a Go Program => 
// $ cd 2-greeting-people
// $ go run greeting.go