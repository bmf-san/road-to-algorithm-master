package main

import (
	"fmt"
)

func linearSearch(a []int, key int) bool {
	for _, item := range a {
		if item == key {
			return true
		}
	}

	return false
}

func main() {
	result := linearSearch([]int{1, 2, 3, 4, 5}, 6)

	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
