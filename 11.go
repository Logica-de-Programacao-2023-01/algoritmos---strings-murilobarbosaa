package main

import (
	"fmt"
	"strings"
)

func removeVowels(str string) string {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vowel := range vowels {
		str = strings.ReplaceAll(str, vowel, "")
	}

	return str
}

func main() {
	str := "Olá, mundo! Esta é uma frase de exemplo."

	result := removeVowels(str)

	fmt.Println("Resultado:", result)
}
