package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	var result = random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("value", value, "is integer")
	default:
	}
}
