package main

import (
	"fmt"
	"unicode"
)

func containsOnlyNumbers(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if containsOnlyNumbers(input) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string contém outros caracteres além de números.")
	}
}
