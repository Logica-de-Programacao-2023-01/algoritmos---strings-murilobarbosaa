package main

import (
	"fmt"
	"strings"
)

func replaceVowelsWithAsterisk(str string) string {
	vowels := "aeiouAEIOU"
	result := ""

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			result += "*"
		} else {
			result += string(char)
		}
	}

	return result
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := replaceVowelsWithAsterisk(input)
	fmt.Println("String com vogais substituídas:", result)
}
