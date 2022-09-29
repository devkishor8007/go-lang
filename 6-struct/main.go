package main

import (
	"fmt" 
	"time"
)

type Food struct {
	Name string
	desc string
	price int
	createAt time.Time
}

func main() {
	food := Food{
		Name: "Momo",
		desc: "It is one of the best food",
		price: 200,
		createAt: time.Now(),
	}

	fmt.Println("Toady I am eating", food.Name, ".", food.desc, "in the market which price is", food.price)
	fmt.Println(food.createAt)
}