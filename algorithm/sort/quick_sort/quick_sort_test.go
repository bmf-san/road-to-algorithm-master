package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := []int{2, 5, 7, 1, 3, 9}

	actual := quickSort(n)
	expected := []int{1, 2, 3, 5, 7, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}
