package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, char, replacement string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&char)

	fmt.Print("Digite o caractere de substituição: ")
	fmt.Scanln(&replacement)

	result := strings.ReplaceAll(input, char, replacement)

	fmt.Println("Resultado:", result)
}
