package main

import "fmt"

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	var fibonacci [5]int = [5]int{1, 1, 2, 3, 5}
	fmt.Println(fibonacci)
	fmt.Println(&fibonacci[0])

	text := [3]string{
		"This is first part of an array",
		"This is the second part",
		"And finally the third part",
	}
	fmt.Println(text[0], text[1], text[2])
}
