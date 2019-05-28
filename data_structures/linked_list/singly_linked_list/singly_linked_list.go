package main

import (
	"errors"
	"fmt"
)

// A node is a node of list.
type node struct {
	value string
	next  *node
}

// A list is a singly linked list.
type list struct {
	head *node
}

// add add a node to tail of a list.
func (l *list) add(newn *node) {
	if l.head == nil {
		l.head = newn
		newn.next = nil
		return
	}

	// sequential access
	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			n.next = newn
			return
		}
	}

	return
}

// insert a node before a particular node of a list.
func (l *list) insert(newn *node, v string) error {
	if l.head == nil {
		return errors.New("a target node is not exists")
	}

	// sequential access
	for n := l.head; n.next != nil; n = n.next {
		if n.next.value == v {
			newn.next = n.next
			n.next = newn
			return nil
		}
	}

	return errors.New("a target node is not exists")
}

// display display all nodes of a list.
func (l *list) display() {
	// sequential access
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.value, n.next)
	}
}

func main() {
	l := &list{}

	first := &node{"first", nil}
	second := &node{"second", nil}
	third := &node{"third", nil}

	l.add(first)
	l.add(second)
	l.add(third)

	between := &node{"between", nil}
	l.insert(between, "second")

	l.display()

	fmt.Printf("%#v\n", l)
}
