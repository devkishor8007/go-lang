package main

import "fmt"

func main() {
	intro := map[string]string{"name": "kishor", "learn": "nodejs", "type": "open soure"}
	fmt.Println(intro)
	fmt.Println(intro["learn"])
}