package main

import (
	"fmt"
	"unicode"
)

func isCamelCase(str string) bool {
	if !unicode.IsUpper(rune(str[0])) {
		return false
	}

	for i := 1; i < len(str)-1; i++ {
		if unicode.IsLower(rune(str[i])) && unicode.IsUpper(rune(str[i+1])) {
			return false
		}
	}

	return true
}

func countWords(str string) int {
	count := 1

	for i := 1; i < len(str); i++ {
		if unicode.IsUpper(rune(str[i])) {
			count++
		}
	}

	return count
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isCamelCase(input) {
		wordCount := countWords(input)
		fmt.Printf("A string está em CamelCase e possui %d palavra(s).\n", wordCount)
	} else {
		fmt.Println("A string não está em CamelCase.")
	}
}
