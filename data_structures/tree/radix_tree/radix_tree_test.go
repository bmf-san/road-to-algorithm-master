/////////////////////////////////////////
// Warning！
// This is an incomplete implementation.
/////////////////////////////////////////

package main

import (
	"reflect"
	"testing"
)

// Insert a null string.
func TestInsertIfWordIsNullString(t *testing.T) {
	cases := []struct {
		actual     *Node
		expected   *Node
		insertItem string
	}{
		{
			actual:     &Node{},
			expected:   &Node{},
			insertItem: "",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestInsertKeyNoneChildrenNone(t *testing.T) {
	cases := []struct {
		actual      *Node
		expected    *Node
		insertItem1 string
	}{
		{
			actual: &Node{},
			expected: &Node{
				key: "foo",
			},
			insertItem1: "foo",
		},
		{
			actual: &Node{},
			expected: &Node{
				key: "bar123",
			},
			insertItem1: "bar123",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem1)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestInsertKeyExistChildrenNoneIfExactMatch(t *testing.T) {
	cases := []struct {
		actual      *Node
		expected    *Node
		insertItem1 string
		insertItem2 string
	}{
		{
			actual: &Node{
				key: "foo",
			},
			expected: &Node{
				key: "foo",
			},
			insertItem1: "foo",
			insertItem2: "foo",
		},
		{
			actual: &Node{
				key: "bar123",
			},
			expected: &Node{
				key: "bar123",
			},
			insertItem1: "bar123",
			insertItem2: "bar123",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem1)
		c.actual.Insert(c.insertItem2)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestInsertKeyExistChildrenNoneIfMisMatch(t *testing.T) {
	cases := []struct {
		actual     *Node
		expected   *Node
		insertItem string
	}{
		{
			actual: &Node{
				key: "foo",
			},
			expected: &Node{
				children: []*Node{
					&Node{
						key: "bar",
					},
					&Node{
						key: "foo",
					},
				},
			},
			insertItem: "bar",
		},
		{
			actual: &Node{
				key: "bar",
			},
			expected: &Node{
				children: []*Node{
					&Node{
						key: "bar",
					},
					&Node{
						key: "foo",
					},
				},
			},
			insertItem: "foo",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestInsertKeyExistChildrenNoneIfSingleNodePrefixMatch(t *testing.T) {
	cases := []struct {
		actual      *Node
		expected    *Node
		insertItem1 string
		insertItem2 string
	}{
		{
			actual: &Node{
				key: "food",
			},
			expected: &Node{
				key: "fo",
				children: []*Node{
					&Node{
						key: "od",
					},
				},
			},
			insertItem1: "fo",
			insertItem2: "food",
		},
		{
			actual: &Node{
				key: "food",
			},
			expected: &Node{
				key: "fo",
				children: []*Node{
					&Node{
						key: "od",
					},
				},
			},
			insertItem1: "food",
			insertItem2: "fo",
		},
		{
			actual: &Node{
				key: "food",
			},
			expected: &Node{
				key: "foo",
				children: []*Node{
					&Node{
						key: "d",
					},
				},
			},
			insertItem1: "foo",
			insertItem2: "food",
		},
		{
			actual: &Node{
				key: "food",
			},
			expected: &Node{
				key: "foo",
				children: []*Node{
					&Node{
						key: "d",
					},
				},
			},
			insertItem1: "food",
			insertItem2: "foo",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem1)
		c.actual.Insert(c.insertItem2)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestInsertKeyExistChildrenNoneIfDoubleNodePrefixMatch(t *testing.T) {
	cases := []struct {
		actual      *Node
		expected    *Node
		insertItem1 string
		insertItem2 string
	}{
		{
			actual: &Node{
				key: "post",
			},
			expected: &Node{
				key: "po",
				children: []*Node{
					&Node{
						key: "em",
					},
					&Node{
						key: "st",
					},
				},
			},
			insertItem1: "post",
			insertItem2: "poem",
		},
		{
			actual: &Node{
				key: "post",
			},
			expected: &Node{
				key: "po",
				children: []*Node{
					&Node{
						key: "em",
					},
					&Node{
						key: "st",
					},
				},
			},
			insertItem1: "poem",
			insertItem2: "post",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem1)
		c.actual.Insert(c.insertItem2)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestInsertKeyNoneChildrenExistIfExactMatch(t *testing.T) {
	cases := []struct {
		actual     *Node
		expected   *Node
		insertItem string
	}{
		{
			actual: &Node{
				children: []*Node{
					&Node{
						key: "bar",
					},
					&Node{
						key: "foo",
					},
				},
			},
			expected: &Node{
				children: []*Node{
					&Node{
						key: "bar",
					},
					&Node{
						key: "foo",
					},
				},
			},
			insertItem: "foo",
		},
		{
			actual: &Node{
				children: []*Node{
					&Node{
						key: "bar",
					},
					&Node{
						key: "foo",
					},
				},
			},
			expected: &Node{
				children: []*Node{
					&Node{
						key: "bar",
					},
					&Node{
						key: "foo",
					},
				},
			},
			insertItem: "bar",
		},
	}

	for _, c := range cases {
		c.actual.Insert(c.insertItem)

		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}
