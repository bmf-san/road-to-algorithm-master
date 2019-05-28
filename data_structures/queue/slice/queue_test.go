package main

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	actual := &queue{}
	actual.enqueue(1)
	actual.enqueue(2)
	actual.enqueue(3)

	expected := &queue{
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

func TestDequeue(t *testing.T) {
	actual := &queue{
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
	actual.dequeue()
	actual.dequeue()
	actual.dequeue()

	expected := &queue{}

	if len(actual.nodes) != len(expected.nodes) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestPeek(t *testing.T) {
	actual := &queue{
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
	actual := &queue{}

	if !reflect.DeepEqual(actual.isEmpty(), true) {
		t.Errorf("actual:%v expected:%v\n", actual.isEmpty(), true)
	}

	actual.enqueue(1)
	if !reflect.DeepEqual(actual.isEmpty(), false) {
		t.Errorf("actual:%v expected:%v\n", actual.isEmpty(), false)
	}
}
