package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	//var year int = now.Year()
	fmt.Println("Today is ", now.Day(), " of month ", now.Month(), " of year ", now.Year())
	fmt.Println(now)
}
