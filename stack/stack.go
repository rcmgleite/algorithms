package stack

import (
	"errors"
	"fmt"
)

// Stacker ...
type Stacker interface {
	Push(int) error
	Pop() (int, error)
	IsEmpty() bool
}

// Print ...
func Print(s Stacker) {
	if s.(*LinkedStack) == nil {
		fmt.Println("Error: Stack is nil")
		return
	}
	for !s.IsEmpty() {
		toPrint, err := s.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(toPrint)
	}
}

// LinkedStack - implemented using a linked list
type LinkedStack struct {
	first *node
	min   *node
	size  int
}

// struct that represents each node of the stack
type node struct {
	item int
	next *node
}

// Size ...
func (s *LinkedStack) Size() int {
	return s.size
}

// IsEmpty returns true if stack is empty
func (s *LinkedStack) IsEmpty() bool {
	return s.first == nil
}

// Min returns the min value on stack
func (s *LinkedStack) Min() int {
	return s.min.item
}

// Push item to stack
func (s *LinkedStack) Push(v int) error {
	newItem := node{item: v, next: s.first}
	s.first = &newItem
	s.size++
	if s.min != nil {
		if s.min.item > s.first.item {
			s.min = s.first
		}
	} else {
		s.min = &newItem
	}
	return nil
}

// Pop item from stack
func (s *LinkedStack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack Underflow")
	}

	toReturn := s.first.item
	s.first = s.first.next
	s.size--
	return toReturn, nil
}

// ArrayStack - implemented using fixed size array
type ArrayStack struct {
	s  []int // stack as the array
	sp int   // stack pointer
}

// NewArrayStack - constructor for ArrayStack
func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{s: make([]int, capacity), sp: -1}
}

// IsEmpty ...
func (s *ArrayStack) IsEmpty() bool {
	return s.sp == -1
}

// Push ...
func (s *ArrayStack) Push(v int) error {
	if s.sp+1 > len(s.s)-1 {
		return errors.New("Stack Overflow")
	}
	s.sp++
	s.s[s.sp] = v
	return nil
}

// Pop ...
func (s *ArrayStack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack Underflow")
	}
	toReturn := s.s[s.sp]
	s.sp--
	return toReturn, nil
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
///////////////////////// 3 stacks with single array ///////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

// Aproach 1: divide the array in equal parts; -- Dummy
const (
	// StackSize ...
	StackSize = 300
)

// TripleStack ...
type TripleStack struct {
	array []int
	sps   [3]int
}

// NewTripleStack ... each stack will have StackSize
func NewTripleStack() *TripleStack {
	array := make([]int, 3*StackSize)
	return &TripleStack{array: array}
}

// Push ...
func (s *TripleStack) Push(value int, stack int) {
	s.sps[stack]++
	arrayIndex := stack*StackSize + s.sps[stack]
	s.array[arrayIndex] = value
}

// Pop ...
// in this function, the pop will NOT clear the array, but the position
// will be free to be overwritten
func (s *TripleStack) Pop(stack int) int {
	arrayIndex := stack*StackSize + s.sps[stack]
	s.sps[stack]--
	return s.array[arrayIndex]
}

// IsEmpty ...
func (s *TripleStack) IsEmpty(stack int) bool {
	return s.sps[stack] == 0
}
