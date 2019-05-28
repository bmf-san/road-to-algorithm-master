package main

import "fmt"

// Node has a key and children.
type Node struct {
	Key      string
	Children map[string]*Node
}

// Tree has a root node.jk
type Tree struct {
	Root *Node
}

// Insert is insert a word to tree.
func (t *Tree) Insert(word string) {
	if t.Root == nil {
		t.Root = &Node{word, make(map[string]*Node)}
		return
	}

	// TODO: write from here...

}

func main() {
	tree := &Tree{}

	// tree.Insert("first")
	// tree.Insert("far")
	// tree.Insert("farm")
	// tree.Insert("feed")
	// tree.Insert("second")

	fmt.Printf("%v", tree)
}
