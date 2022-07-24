package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pemabagi int) (int, error) {
	if pemabagi == 0 {
		return 0, errors.New("Pembagi Tidak Boleh 0")
	} else {
		result := nilai / pemabagi
		return result, nil
	}
}
func main() {
	hasil, err := Pembagi(100, 2)
	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
