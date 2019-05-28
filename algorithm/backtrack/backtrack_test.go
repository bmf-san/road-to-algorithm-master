package main

import (
	"reflect"
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	cases := []struct {
		items []int
		k     int
		rslt  [][]int
	}{
		{
			items: []int{1, 2, 3, 4},
			k:     0,
			rslt: [][]int{
				{},
			},
		},
		{
			items: []int{1, 2, 3, 4},
			k:     1,
			rslt: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			items: []int{1, 2, 3, 4},
			k:     2,
			rslt: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			items: []int{1, 2, 3, 4},
			k:     3,
			rslt: [][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 3, 4},
				{2, 3, 4},
			},
		},
	}

	for _, c := range cases {
		rslt := get(c.items, c.k)
		if reflect.DeepEqual(rslt, c.rslt) == false {
			t.Errorf("actual:%v expected:%v", rslt, c.rslt)
		}
	}
}
