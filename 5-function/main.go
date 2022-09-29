package main

import "fmt"

func main() {
	fmt.Println("function")
	one()
	sum(10, 12)
	form()
}

func one() {
	fmt.Println("Hi, I am ...")
}

func sum(a int, b int) {
	fmt.Println("the sum of a and b is",a + b);
}

func form() {
	fmt.Println("Enter your name")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("My name is %s!. Welcome!!\n", name)
	fmt.Println("Are you sure", name, "well, ...")
}