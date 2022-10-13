package main

import (
	"fmt"
	"time"
)

func f() {
	var i int
	for i = 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, "")
	}
}

func HelloName(name string, age int) {
	fmt.Println(name, age)
}

func main() {
	go f()
	f()
	HelloName("Kishor", 12)
}

