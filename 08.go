package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	chars := strings.Split(str, "")

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	reversed := strings.Join(chars, "")

	return reversed
}

func main() {
	var input string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&input)

	reversed := reverseString(input)

	fmt.Println("String invertida:", reversed)
}
