package main

import (
	"fmt"
	"strings"
)

func replaceLetter(str string, oldChar string, newChar string) string {
	replaced := strings.ReplaceAll(str, oldChar, newChar)

	return replaced
}

func main() {
	var input string
	var oldChar, newChar string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&input)

	fmt.Println("Digite a letra a ser substituída:")
	fmt.Scanln(&oldChar)

	fmt.Println("Digite a letra de substituição:")
	fmt.Scanln(&newChar)

	replaced := replaceLetter(input, oldChar, newChar)

	fmt.Println("String modificada:", replaced)
}
