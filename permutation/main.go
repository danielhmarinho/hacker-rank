package main

import "fmt"

func main() {
	fmt.Println(checkPermutation("teste", "ttest"))
}

// SOLUÇÃO COMPARANDO SE AS DUAS STRINGS SOMANDO E SUBTRAINDO LETRAS POSUEM O VALOR DE 0
// TAMBEM PODERIA SER IMPLEMENTADA USANDO UM SORTE NAS DUAS STRINGS E COMPARANDO SE SÃO IGUAIS (MAIS CUSTOSA)
func checkPermutation(first string, second string) bool {
	charCount := make(map[string]int)
	if len(first) != len(second) {
		return false
	}
	for _, char := range first {
		charCount[string(char)]++
	}

	for _, char := range second {
		charCount[string(char)]--
	}

	for _, value := range charCount {
		if value >= 0 {
			return false
		}
	}
	return true
}
