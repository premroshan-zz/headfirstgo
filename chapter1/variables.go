package main

import (
	"fmt"
	"reflect"
)

func main() {
	quantity := 4
	quantity2 := quantity
	length, width := 1.2, 2.4
	customerName := "Damon Cole"
	quantity = 2
	fmt.Println(float64(quantity))
	quantityFloat := float64(quantity)
	fmt.Println(quantityFloat)
	fmt.Println(reflect.TypeOf(float64(quantity)))
	fmt.Println(customerName, "has ordered", quantity, "sheets each with an area of ", length*width)
	fmt.Println(quantity2)
}
