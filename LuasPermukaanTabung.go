package main

import "fmt"

func main() {
	var r, T, Lp, z float64
	var pi2 float64 = 3.14
	fmt.Print("Masukan Radius Jari-Jari Lingkaran : ")
	fmt.Scanln(&r)
	fmt.Print("Masukan Tinggi Tabung : ")
	fmt.Scanln(&T)
	z = r + T
	Lp = (2 * pi2 * r) * z
	fmt.Println("Hasil Dari Luas Permukaan Tabung Adalah ", Lp)

	fmt.Println("---------------------------------------------------------")
}
