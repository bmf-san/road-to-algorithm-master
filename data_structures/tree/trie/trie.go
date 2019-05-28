package main

import "fmt"

// Node is a node of tree.
type Node struct {
	key      string
	children map[rune]*Node
}

// newTrie create a root node.
func newTrie() *Node {
	return &Node{
		key:      "",
		children: make(map[rune]*Node),
	}
}

// insert insert a word to tree.
func (n *Node) insert(word string) {
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

// search search a word from a tree.
func (n *Node) search(word string) bool {
	if len(n.key) == 0 && len(n.children) == 0 {
		return false
	}

	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			return false
		}
	}

	return true
}

func main() {
	t := newTrie()

	t.insert("word")
	t.insert("wheel")
	t.insert("world")
	t.insert("hospital")
	t.insert("mode")

	fmt.Printf("%v", t.search("mo")) // true
	fmt.Printf("%v", t)
}
