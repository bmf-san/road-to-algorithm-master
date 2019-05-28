package main

import "fmt"

// a binary search tree.
type tree struct {
	root *node
}

// a node for binary search tree.
type node struct {
	val string
	l   *node
	r   *node
}

// insert a value to tree.
func (t *tree) insert(v string) {
	t.root = t.root.insertNode(v)
}

// insert a node to tree.
func (n *node) insertNode(v string) *node {
	if n == nil {
		return &node{val: v}
	}
	if v < n.val {
		n.l = n.l.insertNode(v)
	} else if v > n.val {
		n.r = n.r.insertNode(v)
	}
	return n
}

// search a value from tree.
func (t *tree) search(v string) bool {
	return t.root.searchNode(v)
}

// search a node from tree.
func (n *node) searchNode(v string) bool {
	if n == nil {
		return false
	}
	if v == n.val {
		return true
	} else if v < n.val {
		return n.l.searchNode(v)
	} else {
		return n.r.searchNode(v)
	}
}

// remove a value from tree.
func (t *tree) remove(v string) {
	t.root = t.root.removeNode(v)
}

// remove a node from tree.
func (n *node) removeNode(v string) *node {
	if n == nil {
		return nil
	}

	if v < n.val {
		n.l = n.l.removeNode(v)
		return n
	} else if v > n.val {
		n.r = n.r.removeNode(v)
		return n
	} else {
		// node has no children
		if n.l == nil && n.r == nil {
			return nil
		}

		// node has only right child
		if n.l == nil {
			return n.r
		}

		// node has only left child
		if n.r == nil {
			return n.l
		}

		// node has both children
		leftmostrightside := n.r
		for leftmostrightside.l != nil {
			leftmostrightside = leftmostrightside.l
		}
		n.val = leftmostrightside.val
		n.r = n.r.removeNode(n.val)
		return n
	}
}

// breadth first search - preorder
//
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	5 -> 3 -> 1 -> 4 -> 8 -> 6 -> 9
func (t *tree) preorder(n *node, f func(string)) {
	if n != nil {
		// root → left → right
		f(n.val)
		t.preorder(n.l, f)
		t.preorder(n.r, f)
	}
}

// breadth first search - inorder
//
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	1 -> 3 -> 4 -> 5 -> 6 -> 8 -> 9
func (t *tree) inorder(n *node, f func(string)) {
	if n != nil {
		// left → root → right
		t.inorder(n.l, f)
		f(n.val)
		t.inorder(n.r, f)
	}
}

// breadth first search - postorder
//
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	1 -> 4 -> 3 -> 6 -> 9 -> 8 -> 5
func (t *tree) postorder(n *node, f func(string)) {
	if n != nil {
		// left → right → root
		t.postorder(n.l, f)
		t.postorder(n.r, f)
		f(n.val)
	}
}

// depth first search
//
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	5 -> 3 -> 8 -> 1 -> 4 -> 6 -> 9
func (t *tree) dfs(n *node, f func(string)) {
	if n != nil {
		s := []*node{n}
		for len(s) > 0 {
			crtn := s[0]
			f(crtn.val)
			s = s[1:]
			if crtn.l != nil {
				s = append(s, crtn.l)
			}
			if crtn.r != nil {
				s = append(s, crtn.r)
			}
		}
	}
}

func main() {
	t := &tree{}
	t.insert("5")
	t.insert("3")
	t.insert("8")
	t.insert("1")
	t.insert("4")
	t.insert("6")
	t.insert("9")
	t.insert("11")
	fmt.Println(t.search("11")) // true
	t.remove("11")
	fmt.Println(t.search("11")) // false
	f := func(v string) {
		fmt.Println(v)
	}
	t.preorder(t.root, f) // 5314869
	fmt.Println("-----")
	t.inorder(t.root, f) // 1345689
	fmt.Println("-----")
	t.postorder(t.root, f) // 1436985
	fmt.Println("-----")
	t.dfs(t.root, f) // 5381469
}
