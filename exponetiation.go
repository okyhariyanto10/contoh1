package main

import "fmt"

func main() {

	var i, x, n, power int
	power = 1

	fmt.Print("\nEnter the Number to find the Power = ")
	fmt.Scanln(&x)

	fmt.Print("Enter the nnent Value = ")
	fmt.Scanln(&n)

	for i = 1; i <= n; i++ {
		power = power * x
	}

	fmt.Println(x, " Power ", n, " = ", power)
}
