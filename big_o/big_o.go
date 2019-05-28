// 世界で闘うプログラミング力を鍛える本 VI ビッグ・Ｏ記法（Big O）の例題と練習問題より抜粋したものをGoで実装
// See: https://www.amazon.co.jp/%E4%B8%96%E7%95%8C%E3%81%A7%E9%97%98%E3%81%86%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E5%8A%9B%E3%82%92%E9%8D%9B%E3%81%88%E3%82%8B%E6%9C%AC-%E3%82%B3%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0%E9%9D%A2%E6%8E%A5189%E5%95%8F%E3%81%A8%E3%81%9D%E3%81%AE%E8%A7%A3%E6%B3%95-Gayle-Laakmann-McDowell/dp/4839960100
package main

import (
	"fmt"
	"math"
)

// O(1)
func o1(a int, b int) int {
	if b <= 0 {
		return -1
	}
	div := a / b
	return a - div*b
}

// O(N)
func on_a(items []int) {
	sum := 0
	item := 1
	for i := 0; i < len(items); i++ {
		sum += items[i]
	}
	for j := 0; j < len(items); j++ {
		item *= items[j]
	}
	fmt.Println(sum, item)
}

// O(N)
func on_b(n int) int {
	/// n, n-1, n-2 ...
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		return n * on_b(n-1)
	}
}

// O(N)
func on_c(a int, b int) int {
	sum := 0
	for i := 0; i < b; i++ {
		sum += a
	}
	return sum
}

// O(N)
func on_d(a int, b int) int {
	if b < 0 {
		return 0
	} else if b == 0 {
		return 1
	} else {
		return a * on_d(a, b-1)
	}
}

// O(N^2)
func on2_a(items []int) {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items); j++ {
			fmt.Println(items[i], items[j])
		}
	}
}

// O(N^2)
func on2_b(items []int) {
	// N=5
	// i 0 j 1 2 3 4
	// i 1 j 2 3 4
	// i 2 j 3 4
	// i 3 j 4
	// Think how many times each of the outer loop and the inner loop will loop.
	// The number of loops of j changes depending on the value of i.
	// Considering the average number of loops for j, 10/2 equals 5.
	// N*(N/2) → N^2/2 → O(N^2)

	// loop N times
	for i := 0; i < len(items); i++ {
		// loop N/2 times on average
		for j := i + 1; j < len(items); j++ {
			fmt.Println(items[i], items[j])
		}
	}
}

// O(NM)
func onm_a(a []int, b []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}
}

// O(NM)
func onm_b(a []int, b []int) {
	// N times
	for i := 0; i < len(a); i++ {
		// M times
		for j := 0; j < len(b); j++ {
			// 100 times. 100 is a constant so it doesn't affect the big O.
			for k := 0; k < 100; k++ {
				fmt.Println(a[i], b[j])
			}
		}
	}
}

// O(√N)
func osqrt_a(n int) bool {
	// The loop ends when i reaches the square root of n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// O(√N)
func osqrt_b(n int) int {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	return -1
}

// O(logN)
func ologn_a(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		prev := ologn_a(n / 2)
		curr := prev * 2
		return curr
	}
}

// O(logN)
func ologn_b(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// O(N/M)
func onm(a int, b int) int {
	count := 0
	sum := b
	for sum <= a {
		sum += b
		count++
	}
	return count
}

func main() {
	// O(1)
	o1(10, 3)

	// O(N)
	on_a([]int{1, 2, 3, 4, 5})
	on_b(5)
	on_c(10, 5)
	on_d(10, 5)

	// O(N^2)
	on2_a([]int{1, 2, 3, 4, 5})
	on2_b([]int{1, 2, 3, 4, 5})

	// O(NM)
	onm_a([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
	onm_b([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})

	// O(√N)
	osqrt_a(100)
	osqrt_b(100)

	// O(logN)
	ologn_a(100)
	ologn_b(100)

	// O(N/M)
	onm(10, 5)
}
