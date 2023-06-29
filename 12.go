package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	str = removeSpaces(str)

	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func removeSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func main() {
	str := "radar"

	if isPalindrome(str) {
		fmt.Println("A string", str, "é um palíndromo.")
	} else {
		fmt.Println("A string", str, "não é um palíndromo.")
	}
}
