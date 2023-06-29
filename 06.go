package main

import (
	"fmt"
	"strings"
)

func countWords(str string) int {
	words := strings.Fields(str)

	return len(words)
}

func main() {
	var input string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&input)

	wordCount := countWords(input)

	fmt.Printf("A frase possui %d palavras.\n", wordCount)
}


