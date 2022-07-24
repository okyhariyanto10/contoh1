package main

import "fmt"

func main() {
	var studentScore int
	fmt.Println("MasukaN Nilai Siswa : ")
	fmt.Scanln(&studentScore)
	fmt.Println("==================================================")
	switch {
	case (studentScore <= 100) && (studentScore >= 80):
		fmt.Println("Anda Mendapatkan Nilai A")
	case (studentScore <= 79) && (studentScore >= 65):
		fmt.Println("Anda Mendapatkan Nilai B")
	case (studentScore <= 64) && (studentScore >= 50):
		fmt.Println("Anda Mendapatkan Nilai C")
	case (studentScore <= 49) && (studentScore >= 35):
		fmt.Println("Anda Mendapatkan Nilai D")
	case (studentScore <= 34) && (studentScore >= 0):
		fmt.Println("Anda Mendapatkan Nilai E")
	default:
		{
			fmt.Println("Anda Salah Memasukan Nilai")
		}

	}
}
