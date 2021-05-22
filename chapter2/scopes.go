package main

import "fmt"

var packageVar string = "package"

func main() {
	var functionVar string = "function"

	if true {
		var conditionVar = "condition"
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionVar)
	}
	fmt.Println(packageVar)
	fmt.Println(functionVar)
	//fmt.Println(conditionVar) out of scope
}
