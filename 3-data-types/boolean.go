package main

import "fmt"

func main() {
	isAge := 10 > 1
	fmt.Println(isAge)

	// example: boolean with Flow Control
	book := 200

	if book > 250 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}


// greater than
// 250 > 200 true
// 20 > 50 false

// less than
// 200 < 400 true
// 4 < 2 false

// equal
// 51 = 51 true
// 500 = 280 false