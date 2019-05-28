package main

import (
	"reflect"
	"testing"
)

func TestNewTrie(t *testing.T) {
	actual := newTrie()
	expected := &Node{
		key:      "",
		children: make(map[rune]*Node),
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestSearch(t *testing.T) {
	trie := newTrie()
	trie.insert("word")
	trie.insert("wheel")
	trie.insert("world")
	trie.insert("hospital")
	trie.insert("mode")

	cases := []struct {
		searchItem string
		expected   bool
	}{
		{
			searchItem: "word",
			expected:   true,
		},
		{
			searchItem: "wo",
			expected:   true,
		},
		{
			searchItem: "wh",
			expected:   true,
		},
		{
			searchItem: "wor",
			expected:   true,
		},
		{
			searchItem: "host",
			expected:   false,
		},
		{
			searchItem: "mode",
			expected:   true,
		},
		{
			searchItem: "code",
			expected:   false,
		},
	}

	for _, c := range cases {
		if trie.search(c.searchItem) != c.expected {
			t.Errorf("actual:%v expected:%v\n", c.searchItem, c.expected)
		}
	}
}
