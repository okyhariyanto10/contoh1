package main

import "fmt"

func kosong(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = kosong(3)
	fmt.Println(data)
}
