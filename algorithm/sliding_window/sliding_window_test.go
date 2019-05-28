package main

import (
	"reflect"
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	cases := []struct {
		s    []int
		sum  int
		rslt [][]int
	}{
		{
			s:   []int{1, 3, 2, 6, 4, 9, 9, 5},
			sum: 9,
			rslt: [][]int{
				{1, 3, 2, 6},
				{3, 2, 6},
				{2, 6, 4},
				{6, 4},
				{4, 9},
				{9},
				{9},
			},
		},
		{
			s:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			sum: 5,
			rslt: [][]int{
				{1, 2, 3},
				{2, 3},
				{3, 4},
				{4, 5},
				{5},
				{6},
				{7},
				{8},
				{9},
			},
		},
	}

	for _, c := range cases {
		rslt := SlidingWindow(c.s, c.sum)
		if reflect.DeepEqual(rslt, c.rslt) == false {
			t.Errorf("actual:%v expected:%v", rslt, c.rslt)
		}
	}
}
