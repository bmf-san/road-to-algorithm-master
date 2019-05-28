package main

import (
	"reflect"
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	cases := []struct {
		s      []int
		maxVal int
		rslt   []int
	}{
		{
			s:      []int{5, 3, 3, 4, 1, 2, 2, 3},
			maxVal: 9,
			rslt:   []int{1, 2, 2, 3, 3, 3, 4, 5},
		},
	}

	for _, c := range cases {
		rslt := countSort(c.s, c.maxVal)
		if reflect.DeepEqual(rslt, c.rslt) == false {
			t.Errorf("actual:%v expected:%v", rslt, c.rslt)
		}
	}
}
