package main

import "fmt"

func primeNumber(number int) bool {

	if number%2 == 0 {
		return false
	} else {
		return true

	}
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
