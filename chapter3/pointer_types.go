package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myFloat float64
	var myBool bool
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(reflect.TypeOf(&myBool))
	myBoolPoint := &myBool
	fmt.Println(myBoolPoint)
}
