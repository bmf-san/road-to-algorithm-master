package main

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	s := []int{6, 5, 3, 7, 2, 8}
	h := newHeap(len(s))
	for i := 0; i < len(s); i++ {
		h.insert(s[i])
	}
	h.buildMinHeap()

	expected := []int{2, 3, 5, 7, 6, 8}

	if !reflect.DeepEqual(h.values, expected) {
		t.Errorf("h.values:%v", h.values)
	}

	if h.size != 6 {
		t.Errorf("h.vales:%v", h.size)
	}

	if h.maxsize != 6 {
		t.Errorf("h.maxsize:%v", h.maxsize)
	}
}
