package main

import "fmt"

// Heap is a heap.
type Heap struct {
	values  []int
	size    int
	maxsize int
}

// newHeap creates a heap.
func newHeap(maxsize int) *Heap {
	return &Heap{
		values:  []int{},
		size:    0,
		maxsize: maxsize,
	}
}

// leaf checks whether index is a leaf.
func (h *Heap) leaf(index int) bool {
	if index >= (h.size/2) && index <= h.size {
		return true
	}
	return false
}

// parent checks whether index is a parent.
func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

// leftchild checks whether index is a leftchild.
func (h *Heap) leftchild(index int) int {
	return 2*index + 1
}

// rightchild checks whether index is a rightchild.
func (h *Heap) rightchild(index int) int {
	return 2*index + 2
}

// insert inserts a item to a heap.
func (h *Heap) insert(item int) error {
	if h.size >= h.maxsize {
		return fmt.Errorf("Heal is ful")
	}
	h.values = append(h.values, item)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

// swap swaps two values.
func (h *Heap) swap(first, second int) {
	temp := h.values[first]
	h.values[first] = h.values[second]
	h.values[second] = temp
}

// upHeapify reconstruct a heap for up.
func (h *Heap) upHeapify(index int) {
	for h.values[index] < h.values[h.parent(index)] {
		h.swap(index, h.parent(index))
	}
}

// downHeapify reconstruct a heap for down.
func (h *Heap) downHeapify(current int) {
	if h.leaf(current) {
		return
	}

	smallest := current
	leftChildIndex := h.leftchild(current)
	rightRightIndex := h.rightchild(current)

	if leftChildIndex < h.size && h.values[leftChildIndex] < h.values[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < h.size && h.values[rightRightIndex] < h.values[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		h.swap(current, smallest)
		h.downHeapify(smallest)
	}
	return
}

// buildMinHeap builds a min heap.
func (h *Heap) buildMinHeap() {
	for index := ((h.size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index)
	}
}

// remove removes a value.
func (h *Heap) remove() int {
	top := h.values[0]
	h.values[0] = h.values[h.size-1]
	h.values = h.values[:(h.size)-1]
	h.size--
	h.downHeapify(0)
	return top
}

func heapSort(n []int) []int {
	h := newHeap(len(n))
	for i := 0; i < len(n); i++ {
		h.insert(n[i])
	}
	h.buildMinHeap()

	var r []int
	for i := 0; i < len(n); i++ {
		r = append(r, h.remove())
	}

	return r
}

func main() {
	n := []int{2, 5, 7, 1, 3, 9}
	fmt.Println(heapSort(n))
}
