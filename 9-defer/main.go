package main

import "fmt"

func main() {
	defer fmt.Println("hello")
	defer fmt.Println("hii")
	fmt.Println("first")
	defer fmt.Println("two")
}