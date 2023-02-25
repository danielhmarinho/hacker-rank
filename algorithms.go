package main

import "fmt"

func main() {
	insertionSort([]int{2, 8, 5, 7, 19, 5, 12, 98, 53})
	bubbleSort([]int{2, 8, 5, 7, 19, 5, 12, 98, 53})
}

func insertionSort(n []int) {
	for i := 1; i < len(n); i++ {
		for j := 0; j < i; j++ {
			if n[j] > n[i] {
				n[j], n[i] = n[i], n[j]
			}
		}
	}
	fmt.Println(n)
}

func bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			fmt.Println(i, j)
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
