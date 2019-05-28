package main

import "fmt"

func bubbleSort(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		for j := 0; j < len(n)-i-1; j++ {
			// Compare adjacent values
			if n[j] > n[j+1] {
				// Swap adjacent values
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	return n
}

func main() {
	n := []int{2, 5, 7, 1, 3, 9}
	fmt.Println(bubbleSort(n))
}
