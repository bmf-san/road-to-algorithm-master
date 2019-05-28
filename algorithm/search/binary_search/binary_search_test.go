package main

import "testing"

func TestBinarySearchIterative(t *testing.T) {
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

	var slice []int = []int{1, 3, 5}
	for _, c := range cases {
		result := binarySearchIterative(slice, c.key)
		if result != c.result {
			t.Errorf("actual:%v expected:%v", result, c.result)
		}
	}
}

func TestBinarySearchRecursive(t *testing.T) {
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

	var slice []int = []int{1, 3, 5}
	for _, c := range cases {
		result := binarySearchRecursive(slice, c.key, 0, len(slice)-1)
		if result != c.result {
			t.Errorf("actual:%v expected:%v", result, c.result)
		}
	}
}
