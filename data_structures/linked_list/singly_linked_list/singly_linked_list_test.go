package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		newn1    *node
		newn2    *node
		expected *list
	}{
		{
			newn1: &node{
				"foo",
				nil,
			},
			newn2: &node{
				"bar",
				nil,
			},
			expected: &list{
				head: &node{
					value: "foo",
					next: &node{
						value: "bar",
						next:  nil,
					},
				},
			},
		},
		{
			newn1: &node{
				"",
				nil,
			},
			newn2: &node{
				"foo",
				nil,
			},
			expected: &list{
				head: &node{
					value: "",
					next: &node{
						value: "foo",
						next:  nil,
					},
				},
			},
		},
		{
			newn1: &node{
				"",
				nil,
			},
			newn2: &node{
				"",
				nil,
			},
			expected: &list{
				head: &node{
					value: "",
					next: &node{
						value: "",
						next:  nil,
					},
				},
			},
		},
		{
			newn1: &node{
				"foo",
				nil,
			},
			newn2: &node{
				"",
				nil,
			},
			expected: &list{
				head: &node{
					value: "foo",
					next: &node{
						value: "",
						next:  nil,
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual := &list{}

		actual.add(c.newn1)
		actual.add(c.newn2)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", actual, c.expected)
		}
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		newn     *node
		v        string
		expected *list
	}{
		{
			newn: &node{
				value: "betweenfooandbar",
				next:  nil,
			},
			v: "bar",
			expected: &list{
				head: &node{
					value: "foo",
					next: &node{
						value: "betweenfooandbar",
						next: &node{
							value: "bar",
							next:  nil,
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual := &list{
			head: &node{
				value: "foo",
				next: &node{
					value: "bar",
					next:  nil,
				},
			},
		}

		actual.insert(c.newn, c.v)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", actual, c.expected)
		}
	}
}
