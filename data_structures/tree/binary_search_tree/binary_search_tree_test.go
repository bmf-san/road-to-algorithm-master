package main

import (
	"testing"
)

func newTree() *Tree {
	tree := &Tree{}

	tree.Insert(10)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(15)
	tree.Insert(14)
	tree.Insert(18)
	tree.Insert(16)
	tree.Insert(16)

	return tree
}

func TestInsert(t *testing.T) {
	tree := newTree()

	actual := tree.JSONStringify()
	expected := `{
 "Root": {
  "Key": 10,
  "Left": {
   "Key": 2,
   "Left": null,
   "Right": {
    "Key": 3,
    "Left": null,
    "Right": {
     "Key": 3,
     "Left": null,
     "Right": {
      "Key": 3,
      "Left": null,
      "Right": null
     }
    }
   }
  },
  "Right": {
   "Key": 15,
   "Left": {
    "Key": 14,
    "Left": null,
    "Right": null
   },
   "Right": {
    "Key": 18,
    "Left": {
     "Key": 16,
     "Left": null,
     "Right": {
      "Key": 16,
      "Left": null,
      "Right": null
     }
    },
    "Right": null
   }
  }
 }
}`

	if actual != expected {
		t.Error("Json data is not equal")
	}

	return
}

func TestRemove(t *testing.T) {
	tree := newTree()

	tree.Remove(3)
	tree.Remove(10)
	tree.Remove(16)

	actual := tree.JSONStringify()

	expected := `{
 "Root": {
  "Key": 14,
  "Left": {
   "Key": 2,
   "Left": null,
   "Right": {
    "Key": 3,
    "Left": null,
    "Right": {
     "Key": 3,
     "Left": null,
     "Right": null
    }
   }
  },
  "Right": {
   "Key": 15,
   "Left": null,
   "Right": {
    "Key": 18,
    "Left": {
     "Key": 16,
     "Left": null,
     "Right": null
    },
    "Right": null
   }
  }
 }
}`

	if actual != expected {
		t.Error(("Json data is not equal"))
	}

	return
}

func TestSearch(t *testing.T) {
	tree := newTree()

	actualForExist := tree.Search(10)
	actualForNotExist := tree.Search(19)

	expectedForExist := "Key is exist in tree."
	expectedForNotExist := "Key is not exist in tree."

	if actualForExist != expectedForExist {
		t.Error("Could not found a key.")
	}

	if actualForNotExist != expectedForNotExist {
		t.Error("Could not found a key.")
	}
}
