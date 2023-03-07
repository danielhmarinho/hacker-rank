package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(urlify("Mr John Smith    ", 13))
}

func urlify(url string, size int) string {
	initialSize := len(url)
	finalString := []string{}
	if initialSize < size {
		for i := 0; i <= size-initialSize; i++ {
			url += " "
		}
	}
	for j := 0; j < size; j++ {
		if string(url[j]) == " " {
			finalString = append(finalString, "%20")
		} else {
			finalString = append(finalString, string(url[j]))
		}
	}
	return strings.Join(finalString, "")
}
