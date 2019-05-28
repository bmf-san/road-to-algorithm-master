// cf. https://github.com/TheAlgorithms/Go/blob/master/sorts/merge_sort.go
package main

func merge(a []int, b []int) []int {
	r := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func mergeSort(n []int) []int {
	if len(n) < 2 {
		return n
	}

	var middle = len(n) / 2
	a := mergeSort(n[:middle])
	b := mergeSort(n[middle:])
	return merge(a, b)
}
