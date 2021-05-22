package main

import (
	//	"bufio"
	"fmt"
	//	"os"
	//	"strconv"
)

func main() {
	fmt.Print("Enter a grade:")
	//	reader := bufio.NewReader(os.Stdin)
	//	input, err := reader.ReadString('\n')
	//	grade_integer, err := strconv.Atoi(input)
	var grade_integer int
	fmt.Scanf("%d", &grade_integer)
	fmt.Println("You entered the grade ", grade_integer)
	if grade_integer == 100 {
		fmt.Println("You topped the exam")
	} else if grade_integer >= 60 && grade_integer <= 80 {
		fmt.Println("You get distinction")
	} else if grade_integer >= 40 && grade_integer <= 60 {
		fmt.Println("You passed")
	} else {
		fmt.Println("Oops, you have failed")
	}
}
