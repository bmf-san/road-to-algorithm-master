package main

import (
	"fmt"
)

func binarySearchIterative(a []int, key int) bool {
	start := 0
	end := len(a) - 1
	var i int
	for {
		if end < start {
			return false
		}
		i = (start + end) / 2

		if a[i] == key {
			return true
		}

		if a[i] < key {
			start = i + 1
		} else {
			end = i - 1
		}
	}
}

func binarySearchRecursive(a []int, key int, left int, right int) bool {
	if left > right {
		return false
	}

	var mid int = (left + right) / 2
	if a[mid] == key {
		return true
	} else if key < a[mid] {
		return binarySearchRecursive(a, key, left, mid-1)
	} else {
		return binarySearchRecursive(a, key, mid+1, right)
	}
}

func main() {
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearchIterative(slice, 2))

	fmt.Println(binarySearchRecursive(slice, 2, 0, len(slice)-1))
}
