package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Println("Digite a primeira string:")
	fmt.Scanln(&str1)

	fmt.Println("Digite a segunda string:")
	fmt.Scanln(&str2)

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}
}
