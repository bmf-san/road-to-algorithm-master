package main

import "fmt"

// Node is a node of tree.
type Node struct {
	key      string
	children map[rune]*Node
}

// NewNode is create a root node.
func NewNode() *Node {
	return &Node{
		key:      "",
		children: make(map[rune]*Node),
	}
}

// Insert is insert a word to tree.
func (n *Node) Insert(word string) {
	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			curNode.children[r] = &Node{
				key:      string(r),
				children: make(map[rune]*Node),
			}

			curNode = curNode.children[r]
		}
	}
}

func main() {
	n := NewNode()

	n.Insert("word")
	n.Insert("wheel")
	n.Insert("world")
	n.Insert("hoge")
	n.Insert("hostpital")
	n.Insert("mod")

	fmt.Printf("%v", n)
}
