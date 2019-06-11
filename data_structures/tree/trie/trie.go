package main

import "fmt"

// Node has a key and Children.
type Node struct {
	Key      string
	Children map[rune]*Node
}

// Tree has a root node.
type Tree struct {
	Root *Node
}

// Insert is insert a word to tree.
func (t *Tree) Insert(word string) {
	if t.Root == nil {
		t.Root = &Node{"", make(map[rune]*Node)}
	}

	targetNode := t.Root
	runes := []rune(word)

	for i := range runes {
		r := runes[i]

		if node, ok := targetNode.Children[r]; ok {
			targetNode = node
		} else {
			targetNode.Children[r] = &Node{string(r), make(map[rune]*Node)}

			targetNode = targetNode.Children[r]
		}
	}
}

// Has is check whether tree has a word.
func (t *Tree) Has(word string) bool {
	check := true

	targetNode := t.Root

	runes := []rune(word)

	for i := range runes {
		r := runes[i]

		if node, ok := targetNode.Children[r]; ok {
			targetNode = node
		} else {
			check = false
		}
	}

	return check
}

func main() {
	tree := &Tree{}

	tree.Insert("first")
	tree.Insert("far")
	tree.Insert("farm")
	tree.Insert("feed")
	tree.Insert("second")

	fmt.Printf("%v", tree)

	if tree.Has("far") {
		fmt.Println("A word exists in tree.")
	} else {
		fmt.Println("A word doesn't exists in tree.")
	}
}
