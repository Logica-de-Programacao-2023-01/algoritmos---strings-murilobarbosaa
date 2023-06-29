package main

import (
	"fmt"
	"sort"
)

func sortString(str string) string {
	strBytes := []byte(str)
	sort.Slice(strBytes, func(i, j int) bool {
		return strBytes[i] < strBytes[j]
	})

	return string(strBytes)
}

func main() {
	str1 := "listen"
	str2 := "silent"

	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	if sortedStr1 == sortedStr2 {
		fmt.Println("As strings", str1, "e", str2, "são anagramas.")
	} else {
		fmt.Println("As strings", str1, "e", str2, "não são anagramas.")
	}
}
