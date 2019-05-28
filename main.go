package main

import "fmt"

func countWays(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		a := countWays(n - 1)
		b := countWays(n - 2)
		c := countWays(n - 3)
		return a + b + c
	}
}

func main() {
	fmt.Println(countWays(3))
}
