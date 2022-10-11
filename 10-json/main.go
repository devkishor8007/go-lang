package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("hello json")
	// Using a Map to Generate JSON
	// To encode JSON data we use the Marshal function.

	data := map[string]interface{}{
		"price": 256,
		"name": "GO",
		"isLoading": false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err);
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}