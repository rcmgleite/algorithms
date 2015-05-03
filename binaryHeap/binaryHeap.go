package binaryHeap

import "fmt"

// Based on binary trees but using an array as representation
// 'binaryHeaps.png'

// Insert and DeleteMax -> O(log(n))

// MyHeap ...
type MyHeap struct {
	a        []int
	size     int
	capacity int
}

// NewHeap ... heap constructor
func NewHeap(capacity int) *MyHeap {
	return &MyHeap{a: make([]int, capacity), size: 0, capacity: capacity}
}

// PrintArray ...
func (h *MyHeap) PrintArray() {
	fmt.Println(h.a)
}

// Sort ... At the end, the array of the heap will be sorted
func (h *MyHeap) Sort() {
	for k := h.size / 2; k >= 1; k-- { // construct the heap propertly
		h.Sink(k)
	}
	// swap root with lowest , remove the 'new lowest'(higher value) and sink the 'new root'(lower value)
	for h.size > 1 {
		swap(h.a, 1, h.size)
		h.size--
		h.Sink(1)
	}
}

// Swim will place the node in the correct position on the tree
// by swapping father and son until no violation of heap is found
// The node come from bottom to top with this function
func (h *MyHeap) Swim(k int) {
	for k > 1 && h.a[k/2] < h.a[k] {
		swap(h.a, k, k/2)
		k = k / 2
	}
}

// Sink ... will be used for the function deleteMax
// The node come from top to bottom with this function.
// This means that the node will go down on the tree, until no violation is found
func (h *MyHeap) Sink(k int) {
	for 2*k <= h.size {
		// h.PrintArray()
		j := 2 * k                           // find first child
		if j < h.size && h.a[j] < h.a[j+1] { // verify if this child is the larger
			j++ //if not, increment j to get the index of the larger
		}
		if h.a[j] < h.a[k] { // if child < father no need to swap anymore

			break
		}
		swap(h.a, k, j) // swap father with larger child
		k = j           // make it run again
	}
}

// DeleteMax ...
// the ideia here is to replace the root with the lowest node on the array
// and than sink the new root node unitl no violation is found
func (h *MyHeap) DeleteMax() int {
	max := h.a[1] // the max value will be always the root of the tree, located ate index 1
	swap(h.a, 1, h.size)
	h.size--
	h.Sink(1)
	return max
}

// Insert new node on heap
func (h *MyHeap) Insert(key int) {
	h.size++
	h.a[h.size] = key // add new node at the end of the tree
	h.Swim(h.size)    // swim it to its rigth position
}

//helper function
func swap(array []int, p, q int) {
	aux := array[p]
	array[p] = array[q]
	array[q] = aux
}
