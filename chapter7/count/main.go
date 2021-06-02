package main

import (
	"fmt"
	"log"

	"github.com/premroshan/headfirstgo/chapter7/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
