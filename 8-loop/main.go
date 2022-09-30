package main 

import "fmt"

func main() {
	// Three-component loop
	sum := 10
	for i := 1; i<5; i++ {
		sum += i
	}

	fmt.Println(sum) // 20 (10+1+2+3+4)


	// while loop
	n := 1
	for n < 5 {
		n *= 3
	}
	
	fmt.Println(n) // 9 (1*3*3)

	// for each range loop
	strings := []string{"one", "two", "three"}

	for i, s := range strings {
		fmt.Println(i+1, s)
	}
}