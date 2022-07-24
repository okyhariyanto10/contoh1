package main

import "fmt"

func getCompleteName() (namaPertama, middleName, lastName string) {
	namaPertama = "Oky"
	middleName = "Nisa"
	lastName = "Hariyanto"
	return
}

func main() {
	namaPertama, middleName, lastName := getCompleteName()
	fmt.Println(namaPertama)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
