package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func (customer Customer) sayhay(name string) {
	fmt.Println("Hello ", name, " My Name Is", customer.Name)
}

func (a Customer) sayhuu() {
	fmt.Println("huuuu From", a.Name)
}

func (b Customer) sayhii() {
	fmt.Println("Hii Alamat Rumah", b.Name, "Adalah", b.Address)
}
func (c Customer) sapa() {
	fmt.Println("Hallo Nama Kucingku", c.Name, "Umur Kucing Ku", c.age, "Tahun")
}

func main() {
	var oky Customer
	oky.Name = "Oky"
	oky.Address = "Kalibata"
	oky.age = 26

	var hewan Customer
	hewan.Name = "Kucing"
	hewan.age = 7

	oky.sayhay("Wildan")
	oky.sayhuu()
	oky.sayhii()
	hewan.sapa()

	// fmt.Println(oky.Name)
	// fmt.Println(oky.Address)
	// fmt.Println(oky.age)

	// hariyanto := Customer{
	// 	Name:    "Hariyanto",
	// 	Address: "Cirebon",
	// 	age:     35,
	// }
	// fmt.Println(hariyanto)
}
