package main

import "fmt"

// FIFO queue by using slice.
type queue struct {
	nodes []*node
}

type node struct {
	val int
}

// Add the item to the end of the queue.
func (q *queue) enqueue(item int) {
	// FIFO
	q.nodes = append(q.nodes, &node{
		val: item,
	})
}

// Remove item from the front of the queue.
func (q *queue) dequeue() {
	// FIFO
	q.nodes = q.nodes[1:]
}

// Returns the front item from the queue.
func (q *queue) peek() *node {
	return q.nodes[0]
}

// Returns true if the queue is empty.
func (q *queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func (q *queue) traverse() {
	for _, n := range q.nodes {
		fmt.Printf("%#v\n", n)
	}
}

func main() {
	q := &queue{
		nodes: []*node{
			&node{
				val: 1,
			},
			&node{
				val: 2,
			},
			&node{
				val: 3,
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
