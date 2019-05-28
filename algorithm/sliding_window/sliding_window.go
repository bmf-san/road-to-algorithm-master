package main

import "fmt"

// This function is a example of sliding window which finds subarray from slices.
//
// 1. Define window size (fixed or variable depending on conditions)
// 2. Position the window from the beginning of the data to the start position
// 3. Check if an element in the window satisfies a condition
// 4. Add the subsequence in the window to the result if it satisfies the condition
// 5. Slide the window one step to the right
// 6. Repeat steps 3 through 5 until the window reaches the end of the data
func SlidingWindow(s []int, sum int) [][]int {
	rslt := [][]int{}
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		windowSum += s[windowEnd]

		// found the subarray
		for windowSum >= sum {
			rslt = append(rslt, s[windowStart:windowEnd+1])
			windowSum -= s[windowStart]
			windowStart++
		}
	}

	return rslt
}

func main() {
	s := []int{1, 3, 2, 6, 4, 9, 9, 5}
	sum := 9
	subAry := SlidingWindow(s, sum)
	for _, sa := range subAry {
		fmt.Println(sa)
	}

	// found the subarray steps
	// first
	// [1 3 2 6] ← The window size that meets the conditions is determined
	// [3 2 6] ←Find a window that meets the conditions within that window size

	// second
	// [2 6 4]
	// [6 4]

	// third
	// [4 9]
	// [9]

	// fourth
	// [9]
}
