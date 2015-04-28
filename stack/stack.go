package stack

import (
	"errors"
	"fmt"
)

// Stacker ...
type Stacker interface {
	Push(string) error
	Pop() (string, error)
	IsEmpty() bool
}

// Print ...
func Print(s Stacker) error {
	for !s.IsEmpty() {
		toPrint, err := s.Pop()
		if err != nil {
			return err
		}
		fmt.Println(toPrint)
	}
	return nil
}

// LinkedStack - implemented using a linked list
type LinkedStack struct {
	first *node
	size  int
}

// struct that represents each node of the stack
type node struct {
	item string
	next *node
}

// IsEmpty returns true if stack is empty
func (s *LinkedStack) IsEmpty() bool {
	return s.first == nil
}

// Push item to stack
func (s *LinkedStack) Push(v string) error {
	newItem := node{item: v, next: s.first}
	s.first = &newItem
	s.size++
	return nil
}

// Pop item from stack
func (s *LinkedStack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("Stack Underflow")
	}

	toReturn := s.first.item
	s.first = s.first.next
	s.size--
	return toReturn, nil
}

// ArrayStack - implemented using fixed size array
type ArrayStack struct {
	s  []string // stack as the array
	sp int      // stack pointer
}

// NewArrayStack - constructor for ArrayStack
func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{s: make([]string, capacity), sp: -1}
}

// IsEmpty ...
func (s *ArrayStack) IsEmpty() bool {
	return s.sp == -1
}

// Push ...
func (s *ArrayStack) Push(v string) error {
	if s.sp+1 > len(s.s)-1 {
		return errors.New("Stack Overflow")
	}
	s.sp++
	s.s[s.sp] = v
	return nil
}

// Pop ...
func (s *ArrayStack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("Stack Underflow")
	}
	toReturn := s.s[s.sp]
	s.sp--
	return toReturn, nil
}
