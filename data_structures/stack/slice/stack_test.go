package main

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	actual := &stack{
		nodes: []*node{
			&node{
				val: 1,
			},
			&node{
				val: 2,
			},
			&node{
				val: 3,
			},
		},
	}
	actual.pop()
	actual.pop()
	actual.pop()

	expected := &stack{}

	if len(actual.nodes) != len(expected.nodes) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestPush(t *testing.T) {
	actual := &stack{}
	actual.push(3)
	actual.push(2)
	actual.push(1)

	expected := &stack{
		nodes: []*node{
			&node{
				val: 1,
			},
			&node{
				val: 2,
			},
			&node{
				val: 3,
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestPeek(t *testing.T) {
	actual := &stack{
		nodes: []*node{
			&node{
				val: 1,
			},
		},
	}

	expected := &node{
		val: 1,
	}

	if !reflect.DeepEqual(actual.peek(), expected) {
		t.Errorf("actual:%v expected:%v\n", actual.peek(), expected)
	}
}

func TestIsEmpty(t *testing.T) {
	actual := &stack{}

	if !reflect.DeepEqual(actual.isEmpty(), true) {
		t.Errorf("actual:%v expected:%v\n", actual.isEmpty(), true)
	}

	actual.push(1)
	if !reflect.DeepEqual(actual.isEmpty(), false) {
		t.Errorf("actual:%v expected:%v\n", actual.isEmpty(), false)
	}
}
