package main 

import (
	"fmt"
	"errors"
)

func main() {
	age := 25
	new := errors.New("some")

	if age > 18 {
		fmt.Println("you are qualified")
	} else {
		fmt.Println(new)
	}
}