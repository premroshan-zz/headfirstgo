package main

import (
	"fmt"
)

func main() {
	var quantity, quantity2 int
	var length, width float64
	var customerName string

	quantity = 4
	quantity2 = quantity
	length, width = 1.2, 2.4
	customerName = "Damon Cole"
	quantity = 2
	fmt.Println(customerName, "has ordered", quantity, "sheets each with an area of ", length*width)
	fmt.Println(quantity2)
}
