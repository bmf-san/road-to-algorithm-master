package main

import (
	"reflect"
	"testing"
)

func TestPut(t *testing.T) {
	h := &HashMap{
		data: make(map[int]string),
	}

	cases := []struct {
		key      int
		value    string
		expected *HashMap
	}{
		{
			key:   1,
			value: "foo",
			expected: &HashMap{
				data: map[int]string{
					1: "foo",
				},
			},
		},
		{
			key:   2,
			value: "bar",
			expected: &HashMap{
				data: map[int]string{
					1: "foo",
					2: "bar",
				},
			},
		},
		{
			key:   3,
			value: "foobar",
			expected: &HashMap{
				data: map[int]string{
					1: "foo",
					2: "bar",
					3: "foobar",
				},
			},
		},
	}

	for _, c := range cases {
		h.put(c.key, c.value)
		if !reflect.DeepEqual(h, c.expected) {
			t.Errorf("key:%v value:%v expected:%v", c.key, c.value, c.expected)
		}
	}
}

func TestGet(t *testing.T) {
	h := &HashMap{
		data: make(map[int]string),
	}

	cases := []struct {
		key      int
		value    string
		expected string
	}{
		{
			key:      1,
			value:    "foo",
			expected: "foo",
		},
		{
			key:      2,
			value:    "bar",
			expected: "bar",
		},
		{
			key:      3,
			value:    "foobar",
			expected: "foobar",
		},
	}

	for _, c := range cases {
		h.put(c.key, c.value)
		if h.get(c.key) != c.expected {
			t.Errorf("key:%v value:%v expected:%v", c.key, c.value, c.expected)
		}
	}
}
