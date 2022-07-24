package main

import "fmt"

func main() {

	for brs := 1; brs <= 5; brs++ {
		for spasi := 5; spasi >= brs; spasi-- {
			fmt.Printf(" ")
		}
		for klm := 1; klm <= brs; klm++ {
			fmt.Printf("*")
		}
		for klm := (brs - 1); klm >= 1; klm-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
