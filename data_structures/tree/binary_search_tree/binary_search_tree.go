package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Node has a key and at most two children node.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Tree has a root node.
type Tree struct {
	Root *Node
}

// Insert is insert a node to targetNode.
func (t *Tree) Insert(key int) {
	if t.Root == nil {
		// Create a root node
		t.Root = &Node{key, nil, nil}
	} else {
		RecursiveInsert(t.Root, &Node{key, nil, nil})
	}
}

// RecursiveInsert is insert a new node to target node.
func RecursiveInsert(targetNode *Node, newNode *Node) {
	// Left child key(targetNode.Key) =< Parent key(newNode.Key) =< Right child key(targetNode.Key)
	if newNode.Key < targetNode.Key {
		if targetNode.Left == nil {
			targetNode.Left = newNode
		} else {
			// Recursive calling
			RecursiveInsert(targetNode.Left, newNode)
		}
	} else { // targetNode.Key < newNode.Key
		if targetNode.Right == nil {
			targetNode.Right = newNode
		} else {
			// Recursive calling
			RecursiveInsert(targetNode.Right, newNode)
		}
	}

}

// Remove is remove a key from tree.
func (t *Tree) Remove(key int) {
	RecursiveRemove(t.Root, key)
}

// RecursiveRemove is remove a key from targetNodes recursively.
func RecursiveRemove(targetNode *Node, key int) *Node {
	if targetNode == nil {
		return nil
	}

	if key < targetNode.Key {
		targetNode.Left = RecursiveRemove(targetNode.Left, key)
		return targetNode
	}

	if key > targetNode.Key {
		targetNode.Right = RecursiveRemove(targetNode.Right, key)
		return targetNode
	}

	if targetNode.Left == nil && targetNode.Right == nil {
		targetNode = nil
		return nil
	}

	if targetNode.Left == nil {
		targetNode = targetNode.Right
		return targetNode
	}

	if targetNode.Right == nil {
		targetNode = targetNode.Left
		return targetNode
	}

	leftNodeOfMostRightNode := targetNode.Right

	for {
		if leftNodeOfMostRightNode != nil && leftNodeOfMostRightNode.Left != nil {
			leftNodeOfMostRightNode = leftNodeOfMostRightNode.Left
		} else {
			break
		}
	}

	targetNode.Key = leftNodeOfMostRightNode.Key
	targetNode.Right = RecursiveRemove(targetNode.Right, targetNode.Key)
	return targetNode
}

// Search is search a key from tree.
func (t *Tree) Search(key int) string {
	result := RecursiveSearch(t.Root, key)

	if result {
		return "Key is exist in tree."
	}

	return "Key is not exist in tree."
}

// RecursiveSearch is search a key from tree recursively.
func RecursiveSearch(targetNode *Node, key int) bool {
	if targetNode == nil {
		return false
	}

	// if key is smaller then target node, next time start searching from left node.
	if key < targetNode.Key {
		return RecursiveSearch(targetNode.Left, key)
	}

	// if key is higher then target node, next time start searching from right node.
	if key > targetNode.Key {
		return RecursiveSearch(targetNode.Right, key)
	}

	return true
}

// InOrderTraverse is traverse by in-order.
func (t *Tree) InOrderTraverse() {
	RecursiveInOrderTravese(t.Root)
}

// RecursiveInOrderTravese is traverse by in-order recursively.
func RecursiveInOrderTravese(n *Node) {
	if n != nil {
		RecursiveInOrderTravese(n.Left)
		fmt.Printf("%d, ", n.Key)
		RecursiveInOrderTravese(n.Right)
	}
}

// PreOrderTraverse is traverse by pre-order.
func (t *Tree) PreOrderTraverse() {
	RecursivePreOrderTraverse(t.Root)
}

// RecursivePreOrderTraverse is travese by pre-order recursively.
func RecursivePreOrderTraverse(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.Key)
		RecursivePreOrderTraverse(n.Left)
		RecursivePreOrderTraverse(n.Right)
	}
}

// PostOrderTraverse is traverse by post-order.
func (t *Tree) PostOrderTraverse() {
	RecursivePostOrderTraverse(t.Root)
}

// RecursivePostOrderTraverse is traverse by post-order recursively.
func RecursivePostOrderTraverse(n *Node) {
	if n != nil {
		RecursivePostOrderTraverse(n.Left)
		RecursivePostOrderTraverse(n.Right)
		fmt.Printf("%d ", n.Key)
	}
}

// LevelOrderTraverse is traverse by level-order.
func (t *Tree) LevelOrderTraverse() {
	if t != nil {
		queue := []*Node{t.Root}

		for len(queue) > 0 {
			currentNode := queue[0]
			fmt.Printf("%d ", currentNode.Key)

			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}
}

// JSONStringify is output tree by json strings.
func (t *Tree) JSONStringify() string {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}

	jsonStr := string(jsonBytes)
	var buf bytes.Buffer

	err = json.Indent(&buf, []byte(jsonStr), "", " ")
	if err != nil {
		return err.Error()
	}

	return buf.String()
}

func main() {
	// Intialize a Tree
	tree := &Tree{}

	// Insert
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

	// Remove
	tree.Remove(3)
	tree.Remove(10)
	tree.Remove(16)

	// Search
	fmt.Println(tree.Search(10))
	fmt.Println(tree.Search(19))

	// Traverse
	tree.InOrderTraverse()
	tree.PreOrderTraverse()
	tree.PostOrderTraverse()
	tree.LevelOrderTraverse()

	// Output
	fmt.Printf("%s", tree.JSONStringify())
}
