// runner technique
package main

import "fmt"

type node struct {
	val  string
	next *node
}

type list struct {
	head *node
}

// find the nth to last element of a singly linked list
func (l *list) search(n int) *node {
	n1 := l.head
	n2 := l.head

	// find the nth node from the end.
	for i := 0; i < n; i++ {
		if n1 == nil {
			return nil
		}
		n1 = n1.next
	}

	// n2 scans from the top node, n1 scans from the kth node.
	// When n1 reaches the tail node, n2 is the nth node from the end of the node.
	for n1 != nil {
		n1 = n1.next
		n2 = n2.next
	}

	return n2
}

func main() {
	l := &list{
		head: &node{
			val: "a",
			next: &node{
				val: "b",
				next: &node{
					val: "c",
					next: &node{
						val: "d",
					},
				},
			},
		},
	}
	fmt.Printf("%+v\n", l.search(1)) // d
	fmt.Printf("%+v\n", l.search(2)) // c
	fmt.Printf("%+v\n", l.search(3)) // b
	fmt.Printf("%+v\n", l.search(4)) // a
}
