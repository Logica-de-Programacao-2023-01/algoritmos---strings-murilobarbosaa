package main

import (
	"fmt"
	"unicode"
)

func checkForNumber(str string) bool {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func main() {
	var input string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&input)

	hasNumber := checkForNumber(input)

	if hasNumber {
		fmt.Println("A frase contém pelo menos um número.")
	} else {
		fmt.Println("A frase não contém nenhum número.")
	}
}
