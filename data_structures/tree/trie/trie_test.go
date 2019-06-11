package main

import (
	"testing"
)

func newTree() *Tree {
	tree := &Tree{}

	tree.Insert("first")
	tree.Insert("far")
	tree.Insert("farm")
	tree.Insert("feed")
	tree.Insert("second")

	return tree
}

func TestHas(t *testing.T) {
	tree := newTree()

	actual := tree.Has("first")
	expected := true

	if actual != expected {
		t.Error("A word doesn't existed in tree.")
	}

	actual = tree.Has("firs")
	expected = true

	if actual != expected {
		t.Error("A word doesn't existed in tree.")
	}

	actual = tree.Has("far")
	expected = true

	if actual != expected {
		t.Error("A word doesn't existed in tree.")
	}

	actual = tree.Has("farmer")
	expected = false

	if actual != expected {
		t.Error("A word doesn't existed in tree.")
	}

	actual = tree.Has("second")
	expected = true

	if actual != expected {
		t.Error("A word doesn't existed in tree.")
	}

	actual = tree.Has("secondd")
	expected = false

	if actual != expected {
		t.Error("A word doesn't existed in tree.")
	}
}
