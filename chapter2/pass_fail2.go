package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	grade_integer, err := strconv.ParseInt(input, 10, 10)
	//fmt.Scanf("%d", &grade_integer)
	if err != nil {
		log.Fatal(err)
	}
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
