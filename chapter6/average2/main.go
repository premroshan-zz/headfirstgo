package main

import (
	"fmt"
	"log"

	"github.com/premroshan/headfirstgo/chapter6/datafile"
)

func main() {
	//numbers := [3]float64{21.54, 21.8, 13.4}
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	arrayCount := len(numbers)
	fmt.Printf("Average is %.2f", sum/float64(arrayCount))
}
