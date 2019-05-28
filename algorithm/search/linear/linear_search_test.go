package main

import "testing"

func TestLinearSearch(t *testing.T) {
	cases := []struct {
		key    int
		result bool
	}{
		{
			key:    1,
			result: true,
		},
		{
			key:    2,
			result: false,
		},
	}

	for _, c := range cases {
		result := linearSearch([]int{1, 3, 5}, c.key)
		if result != c.result {
			t.Errorf("actual:%v expected:%v", result, c.result)
		}
	}
}
