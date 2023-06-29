package main

import (
	"fmt"
)

func findUniqueLetters(str string) string {
	letterCount := make(map[rune]int)

	for _, char := range str {
		letterCount[char]++
	}

	uniqueLetters := ""

	for _, char := range str {
		if letterCount[char] == 1 {
			uniqueLetters += string(char)
		}
	}

	return uniqueLetters
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	uniqueLetters := findUniqueLetters(input)
	fmt.Println("Letras únicas na string:", uniqueLetters)
}
