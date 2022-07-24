package main

import "fmt"

//FAKTORIAL MENGGUNAKAN LOOPING
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//FAKTORIAL MENGGUNAKAN RECURSIVE FUNCTION
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(7)
	fmt.Println(loop)

	recursive := factorialRecursive(7)
	fmt.Println(recursive)

}
