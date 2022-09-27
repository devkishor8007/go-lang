package main

import "fmt"

func main() {
	// Raw String Literals
	name := `hello "kishor", learn go!`
	bio := `I can do and try to work on this,
	hope can be worked as we eed for the porduct`
	fmt.Println(name)
	fmt.Println(bio)

	// Interpreted String Literals

	color := "choose the color \"red\" to go!"
	fmt.Println(color)
}