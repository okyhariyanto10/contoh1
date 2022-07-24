package main

import "fmt"

//hello World
// func main() {
// 	fmt.Println("Hello World")
// }

func main() {
	//CONTOH VARIABEL
	var isLogin bool = false
	fmt.Println(isLogin)

	var prime int = 7
	fmt.Println(prime)

	var decimal float64 = 45.6
	fmt.Println(decimal)

	var hello string = "Hello World"
	fmt.Println(hello)

	const pi = 3.14
	fmt.Println(pi)

	// CONTOH OPERAND DAN OPERATOR
	x := 1 + 2
	fmt.Println(x)

	//EXPRESION
	a := 5
	b := 6
	c := a + b
	fmt.Println(c)

	// CONTOH LUAS SEGITIGA
	var r, T, Lp, z float64
	var pi2 float64 = 3.14
	fmt.Print("Masukan Radius Jari-Jari Lingkaran : ")
	fmt.Scanln(&r)
	fmt.Print("Masukan Tinggi Tabung : ")
	fmt.Scanln(&T)
	z = r + T
	Lp = (2 * pi2 * r) * z
	fmt.Println(Lp)

	//STRING OPERATION
	helloworld := "Hello" + " " + "World"
	fmt.Println(helloworld)

	//IF, ELSE IF, ELSE
	hour := 21
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
	//SWITCH
	var today int = 2
	switch today {
	case 1:
		fmt.Println("Today Is Monday")
	case 2:
		fmt.Println("Today Is Tuesday")
	default:
		fmt.Println("Invalid Day")
	}
	//LOOPING
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}
	fmt.Println("")
	//LOOPING CONTINUE & BREAK
	for x := 0; x < 5; x++ {
		if x == 1 {
			continue
		}
		if x > 3 {
			break
		}
		fmt.Println(x)
	}
	//LOOPING OVER STRING(mengukur panjang string)
	sentence := "Hello"
	for y := 0; y < len(sentence); y++ {
		fmt.Printf(string(sentence[y]) + "-")
	}
	fmt.Println("   ----->>>")
	for pos, char := range sentence {
		fmt.Printf("Character %c Start at byte position %d\n", char, pos)
	}
}
