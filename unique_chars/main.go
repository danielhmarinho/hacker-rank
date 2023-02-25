package main

import "fmt"

func main() {
	fmt.Println(isUnique("woraded"))
}

func isUnique(word string) bool {
	charCount := make(map[string]int)
	for _, char := range word {
		charCount[string(char)]++
		if charCount[string(char)] > 1 {
			return false
		}
	}
	return true
}
