package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for x := 1; x <= number; x++ {
			fmt.Printf("%4d", i*x)

		}
		fmt.Printf("\n")
	}

}
func main() {
	cetakTablePerkalian(6)
}
