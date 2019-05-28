package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	actual := &tree{}
	actual.insert("5")
	actual.insert("3")
	actual.insert("8")
	actual.insert("1")
	actual.insert("4")
	actual.insert("6")
	actual.insert("9")

	expected := &tree{
		root: &node{
			val: "5",
			l: &node{
				val: "3",
				l: &node{
					val: "1",
				},
				r: &node{
					val: "4",
				},
			},
			r: &node{
				val: "8",
				l: &node{
					val: "6",
				},
				r: &node{
					val: "9",
				},
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestSearch(t *testing.T) {
	actual := &tree{}
	actual.insert("5")
	actual.insert("3")
	actual.insert("8")
	actual.insert("1")
	actual.insert("4")
	actual.insert("6")
	actual.insert("9")

	testCases := []struct {
		item     string
		expected bool
	}{
		{
			item:     "5",
			expected: true,
		},
		{
			item:     "1",
			expected: true,
		},
		{
			item:     "9",
			expected: true,
		},
		{
			item:     "10",
			expected: false,
		},
		{
			item:     "0",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.item, func(t *testing.T) {
			if actual := actual.search(tC.item); actual != tC.expected {
				t.Errorf("actual:%v expected:%v", actual, tC.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	actual := &tree{}
	actual.insert("5")
	actual.insert("3")
	actual.insert("8")
	actual.insert("1")
	actual.insert("4")
	actual.insert("6")
	actual.insert("9")

	actual.remove("5")
	actual.remove("3")
	actual.remove("8")
	actual.remove("1")
	actual.remove("4")
	actual.remove("6")
	actual.remove("9")

	expected := &tree{}

	actual.dfs(actual.root, func(v string) {
		fmt.Println(v)
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestPreOrder(t *testing.T) {
	bst := &tree{}
	bst.insert("5")
	bst.insert("3")
	bst.insert("8")
	bst.insert("1")
	bst.insert("4")
	bst.insert("6")
	bst.insert("9")

	expected := []string{"5", "3", "1", "4", "8", "6", "9"}

	var actual []string

	bst.preorder(bst.root, func(v string) {
		actual = append(actual, v)
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestInOrder(t *testing.T) {
	bst := &tree{}
	bst.insert("5")
	bst.insert("3")
	bst.insert("8")
	bst.insert("1")
	bst.insert("4")
	bst.insert("6")
	bst.insert("9")

	expected := []string{"1", "3", "4", "5", "6", "8", "9"}

	var actual []string

	bst.inorder(bst.root, func(v string) {
		actual = append(actual, v)
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestPostOrder(t *testing.T) {
	bst := &tree{}
	bst.insert("5")
	bst.insert("3")
	bst.insert("8")
	bst.insert("1")
	bst.insert("4")
	bst.insert("6")
	bst.insert("9")

	expected := []string{"1", "4", "3", "6", "9", "8", "5"}

	var actual []string

	bst.postorder(bst.root, func(v string) {
		actual = append(actual, v)
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestDfs(t *testing.T) {
	bst := &tree{}
	bst.insert("5")
	bst.insert("3")
	bst.insert("8")
	bst.insert("1")
	bst.insert("4")
	bst.insert("6")
	bst.insert("9")

	expected := []string{"5", "3", "8", "1", "4", "6", "9"}

	var actual []string

	bst.dfs(bst.root, func(v string) {
		actual = append(actual, v)
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}
