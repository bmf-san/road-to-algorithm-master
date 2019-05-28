package main

import "fmt"

func backtrack(rslt *[][]int, tmp []int, items []int, start int, k int) {
	if k == 0 {
		combination := make([]int, len(tmp))
		copy(combination, tmp)
		*rslt = append(*rslt, combination)
		return
	}

	for i := start; i < len(items); i++ {
		tmp = append(tmp, items[i])
		backtrack(rslt, tmp, items, i+1, k-1)
		tmp = tmp[:len(tmp)-1]
	}
}

// get a combination from slice.
func get(items []int, k int) [][]int {
	rslt := [][]int{}
	tmp := []int{}
	backtrack(&rslt, tmp, items, 0, k)
	return rslt
}

// ex. items [1, 2, 3, 4]
// k = 0
// N/A
// k = 1
// 1  2  3  4
// k = 2
//
//	   1      2    3
//	 / | \   / \   |
//	2  3  4 3   4  4
//
// k = 3
//
//	    1     2
//	   / \    |
//	  2   3   3
//	 / \  |   |
//	3  4  4   4
//
// k = 4
// 1
// 2
// 3
// 4
// k is the number of values to extract from the element, but it is also the depth to search the tree.
func main() {
	items := []int{1, 2, 3, 4}
	k := 3
	combination := get(items, k)
	fmt.Println(combination)
}
