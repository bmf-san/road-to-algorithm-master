package main

import "fmt"

func selectionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		min := i

		// Compare the smallest value in the data with the first value
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[min] {
				min = j
			}
		}

		// Swap
		n[i], n[min] = n[min], n[i]
	}

	return n
}

func main() {
	n := []int{2, 5, 7, 1, 3, 9}
	fmt.Println(selectionSort(n))
}
