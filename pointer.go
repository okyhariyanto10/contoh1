package main

import "fmt"

type Tinggal struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(tinggal *Tinggal) {
	tinggal.Country = "Indonesia"
}
func main() {
	var tempatTinggal1 Tinggal = Tinggal{"Jaksel", "DKI Jakarta", "Indonesia"}
	var tempatTinggal2 *Tinggal = &tempatTinggal1
	var tempatTinggal3 *Tinggal = &tempatTinggal1

	tempatTinggal2.City = "Bandung"
	*tempatTinggal2 = Tinggal{"Malang", "Jawa Timur", "Indonesia"}
	*tempatTinggal3 = Tinggal{"Purwokerto", "Jawa Tengah", "Indoensia"} //memakasa mengubah data variable baru(smua data akan ikut berubah diatasnya)

	var tempatTinggal4 *Tinggal = new(Tinggal) //function membuat address kosong
	tempatTinggal4.City = "Bogor"

	var tempatTinggal5 = Tinggal{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Tinggal = &tempatTinggal5
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)

	fmt.Println(tempatTinggal1)
	fmt.Println(tempatTinggal2)
	fmt.Println(tempatTinggal3)

	fmt.Println(tempatTinggal4)
}
