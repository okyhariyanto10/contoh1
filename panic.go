package main

import "fmt"

func endApp() {
	massage := recover()
	if massage != nil {
		fmt.Println("Error Dengan Massage : ", massage)
	}
	fmt.Println("Aplikasi Selesai")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi Berjalan")
}
func main() {
	runApp(true)
}
