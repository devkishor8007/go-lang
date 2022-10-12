package main

import "fmt"

func main() {
	intro := map[string]string{"name": "kishor", "learn": "nodejs", "type": "open soure"}
	fmt.Println(intro)
	fmt.Println(intro["learn"])

	employee := map[string]int{"salary": 1000, "tds": 120, "serviceCharge": 150}
	fmt.Println(employee)
	fmt.Println(employee["tds"])


	// The make function takes as argument the type of the map and it returns an initialized map.
	company := make(map[string]string)
	company["name"] = "tech" // Add element
	company["place"] = "abc" // Add element
	fmt.Println(company)
	fmt.Println(len(company)) 
}