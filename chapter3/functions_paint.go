package main

import "fmt"

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(3.1, 5.0)
}

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f litres of paint needed \n", area/10.0)
}
