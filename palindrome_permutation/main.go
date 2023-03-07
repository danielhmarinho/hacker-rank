package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(palindromePermutation("tacteae Coa"))
}

func palindromePermutation(text string) bool {
	charCount := make(map[string]int)
	oddChars := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			charCount[strings.ToLower(string(char))]++
		}
	}
	for _, count := range charCount {
		if count%2 == 1 {
			oddChars++
		}
	}
	return oddChars <= 1
}
