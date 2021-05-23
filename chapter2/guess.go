package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100. Can you guess it?")
	//fmt.Println(target)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10-guesses, " guesses left")
		fmt.Print("Can you guess the number:")
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		guess_int, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if guess_int < target {
			fmt.Println("Your guess is lower than target")
		} else if guess_int > target {
			fmt.Println("Your guess is greater than the target")
		} else if guess_int == target {
			fmt.Println("Your guess is correct. Thanks for playing")
			break
		}
	}
}
