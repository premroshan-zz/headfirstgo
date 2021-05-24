package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a number:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	input_float, err := strconv.ParseFloat(input, 64)
	root, err := squareRoot(input_float)
	if err != nil {
		log.Fatal(err)
	}
	printRoot(input_float, root)
}

func squareRoot(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("Input cannot be negative")
	} else {
		return math.Sqrt(input), nil
	}
}

func printRoot(input float64, root float64) {
	fmt.Printf("The square root of %.2f is %.2f", input, root)
}
