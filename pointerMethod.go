package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
func main() {
	oky := Man{"Oky"}
	oky.Married()

	fmt.Println(oky.Name)
}
