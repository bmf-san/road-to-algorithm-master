package main

import "fmt"

func countSort(s []int, maxVal int) []int {
	count := make([]int, maxVal+1)

	// Count the number of occurrences by using the index as a number
	for _, num := range s {
		count[num]++
	}

	// Calculate cumulative sum
	// ex. [1, 2, 3, 4, 5] →　[1, 3, 6, 10, 15]
	for i := 1; i <= maxVal; i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(s))

	// Traverse the original array from the end and place the elements into the sorted result
	for i := len(s) - 1; i >= 0; i-- {
		num := s[i]
		count[num]--
		sorted[count[num]] = num
	}

	return sorted
}

func main() {
	s := []int{5, 3, 2, 1, 2, 3, 4}
	maxVal := 5
	r := countSort(s, maxVal)
	fmt.Println(r)
}
