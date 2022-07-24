package main

import (
	"fmt"
)

func palindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {

		if word[i] != word[len(word)-1-i] {
			return false
		}

	}

	return true
}

func main() {
	result := palindrome("katak")
	if result == true {
		fmt.Println("Ini Kata Palindrome")
	} else {
		fmt.Println("Ini Bukan Palindrome")
	}
}
