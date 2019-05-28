package main

import "fmt"

// FIFO queue by using linked.
type queue struct {
	top *node
}

type node struct {
	val  int
	next *node
}

// Add the item to the end of the queue.
func (q *queue) enqueue(item int) {
	// FIFO
	if q.top == nil {
		q.top = &node{
			val: item,
		}
		return
	}

	crt := q.top
	for {
		if crt.next == nil {
			crt.next = &node{
				val: item,
			}
			break
		}
		crt = crt.next
	}
}

// Remove item from the front of the queue.
func (q *queue) dequeue() {
	// FIFO
	q.top = q.top.next
}

// Returns the front item from the queue.
func (q *queue) peek() *node {
	return q.top
}

// Returns true if the queue is empty.
func (q *queue) isEmpty() bool {
	return q.top == nil
}

func (q *queue) traverse() {
	crt := q.top
	for {
		if crt == nil {
			break
		}
		fmt.Printf("%#v\n", crt)
		crt = crt.next
	}
}

func main() {
	q := &queue{
		top: &node{
			val: 1,
			next: &node{
				val: 2,
				next: &node{
					val: 3,
				},
			},
		},
	}
	q.dequeue()
	q.traverse()
	fmt.Println("----") // 2 3
	q.dequeue()
	q.traverse() // 3
	fmt.Println("----")
	q.dequeue()
	q.traverse() // nil

	q.enqueue(1)
	q.traverse() // 1
	fmt.Println("----")
	q.enqueue(2)
	q.enqueue(3)
	q.traverse() // 1 2 3

	fmt.Println("----")
	fmt.Printf("%#v\n", q.peek()) // 3

	q2 := &queue{}
	fmt.Println(q2.isEmpty()) // true
}
