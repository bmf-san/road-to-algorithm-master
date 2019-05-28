package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	n := []int{2, 5, 7, 1, 3, 9}

	actual := bubbleSort(n)
	expected := []int{1, 2, 3, 5, 7, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}
