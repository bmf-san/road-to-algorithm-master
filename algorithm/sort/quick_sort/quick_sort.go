package main

import (
	"fmt"
	"math/rand"
)

func quickSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	pivot := n[rand.Intn(len(n))]

	low := make([]int, 0, len(n))
	high := make([]int, 0, len(n))
	middle := make([]int, 0, len(n))

	for _, i := range n {
		switch {
		case i < pivot:
			low = append(low, i)
		case i == pivot:
			middle = append(middle, i)
		case i > pivot:
			high = append(high, i)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	low = append(low, middle...)
	low = append(low, high...)

	return low
}

func main() {
	n := []int{2, 5, 7, 1, 3, 9}
	fmt.Println(quickSort(n))
}
