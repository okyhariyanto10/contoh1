package main

import "fmt"

func main() {
	var bilangan int
	fmt.Println("Masukan Bilangan Faktorial : ")
	fmt.Scanln(&bilangan)
	fmt.Println("")
	fmt.Println("Bilangan Faktorial nya Adalah ")

	for i := 1; i < bilangan; i++ {
		x := 0
		for y := 1; y < bilangan; y++ {
			if i%y == 0 {
				x++
			}

		}
		if (x == 2) && (i != 1) {
			fmt.Println(i)
		}
	}

}
