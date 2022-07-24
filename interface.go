package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayhello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	NameAnimal string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.NameAnimal
}

func main() {
	var oky Person
	oky.Name = "Oky Hariyanto"

	var singa Animal
	singa.NameAnimal = "Singa"

	sayhello(oky)
	sayhello(singa)

}
