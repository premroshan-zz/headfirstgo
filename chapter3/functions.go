package main

import "fmt"

func main() {
	repeatLine("Hello", 3)
}

func repeatLine(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(text)
	}
}
