package main

import (
	"fmt"

	"github.com/rcmgleite/algorithms/customSort"
)

func main() {
	// v := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	// fmt.Println(binarySearch.Recursive(v, 10))
	// fmt.Println(binarySearch.Iterative(v, 9))

	// s := stack.LinkedStack{}
	// s.Push("rafael")
	// s.Push("Camargo")
	// s.Push("Leite")
	// stack.Print(&s)

	// s2 := stack.NewArrayStack(5)
	// s2.Push("rafael")
	// s2.Push("Camargo")
	// s2.Push("Leite")
	// stack.Print(s2)
	//
	// // Overflow test
	// s2.Push("rafael")
	// s2.Push("rafael")
	// s2.Push("rafael")
	// s2.Push("rafael")
	// s2.Push("rafael")
	// err := s2.Push("rafael")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// stack.Print(s2)
	//
	// // Underflow test
	// poped, err := s2.Pop()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(poped)
	// }

	// q := &queue.LinkedQueue{}
	// q.Enqueue("1")
	// q.Enqueue("2")
	// q.Enqueue("3")
	// q.Enqueue("4")
	// q.Enqueue("5")
	// q.Enqueue("6")
	//
	// for !q.IsEmpty() {
	// 	fmt.Println(q.Dequeue())
	// }

	// v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// customSort.Shuffle(v)
	// fmt.Println(v)
	// fmt.Println(customSort.IsSorted(v))
	// customSort.ByShell(v)
	// fmt.Println(v)
	// fmt.Println(customSort.IsSorted(v))

	// v1 := []int{4, 6, 7, 10, 12, 2, 3, 5, 8, 9}
	v1 := []int{12, 4, 5, 2, 3, 11, 6, 7, 10, 9}
	fmt.Println(v1)
	customSort.MergeSort(v1)
	fmt.Println("Final result ", v1)
	fmt.Printf("arraySorted = %v\n", customSort.IsSorted(v1))

}
