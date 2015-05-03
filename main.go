package main

import (
	"fmt"

	"github.com/rcmgleite/algorithms/binaryHeap"
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

	// v1 := []int{12, 4, 5, 2, 3, 11, 6, 7, 10, 9}
	// fmt.Println(v1)
	// customSort.MergeSort(v1)
	// fmt.Println("Final result ", v1)
	// fmt.Printf("arraySorted = %v\n", customSort.IsSorted(v1))

	// var vertexNum = 6
	// g := graph.NewGraph(vertexNum)
	// g.AddEdge(0, 5)
	// g.AddEdge(2, 4)
	// g.AddEdge(2, 3)
	// g.AddEdge(1, 2)
	// g.AddEdge(0, 1)
	// g.AddEdge(3, 4)
	// g.AddEdge(3, 5)
	// g.AddEdge(0, 2)
	//
	// for i := 0; i < vertexNum; i++ {
	// 	g.PrintAdjList(i)
	// }
	//
	// fmt.Printf("\n\n")
	//
	// bfsPaths := graph.NewBfsPaths(*g, 0)
	// fmt.Println("Marked : ", bfsPaths.Marked)
	// fmt.Println("Edge To : ", bfsPaths.EdgeTo)
	// fmt.Println("Dist To : ", bfsPaths.DistTo)

	// var vertexNum = 13
	// g := graph.NewGraph(vertexNum)
	// g.AddEdge(0, 5)
	// g.AddEdge(4, 3)
	// g.AddEdge(0, 1)
	// g.AddEdge(9, 12)
	// g.AddEdge(6, 4)
	// g.AddEdge(5, 4)
	// g.AddEdge(0, 2)
	// g.AddEdge(11, 12)
	// g.AddEdge(9, 10)
	// g.AddEdge(0, 6)
	// g.AddEdge(7, 8)
	// g.AddEdge(9, 11)
	// g.AddEdge(5, 3)
	//
	// for i := 0; i < vertexNum; i++ {
	// 	g.PrintAdjList(i)
	// }
	//
	// cc := graph.NewCC(g)
	//
	// fmt.Println("Marked = ", cc.Marked)
	// fmt.Println("Ids = ", cc.Ids)
	//
	// fmt.Printf("IsConnected(7, 8) = %v\n", cc.IsConnected(7, 8))
	// fmt.Printf("IsConnected(7, 9) = %v\n", cc.IsConnected(7, 9))
	// fmt.Printf("IsConnected(9, 12) = %v\n", cc.IsConnected(9, 12))
	// fmt.Printf("IsConnected(0, 1) = %v\n", cc.IsConnected(0, 1))
	// fmt.Printf("IsConnected(3, 8) = %v\n", cc.IsConnected(3, 8))
	//
	// fmt.Printf("\n\n")

	// ht := &hashTable.HashTable{}
	// ht.Put(0, 11)
	// ht.Put(128, 22)
	// ht.Put(129, 33)
	// ht.Put(2, 44)
	//
	// ht.PrintAllEntries()
	// fmt.Println("ht.Get(0) = ", ht.Get(0))
	// fmt.Println("ht.Get(1) = ", ht.Get(128))
	// fmt.Println("ht.Get(2) = ", ht.Get(129))
	// fmt.Println("ht.Get(3) = ", ht.Get(2))

	// v1 := []int{12, 4, 5, 2, 3, 11, 6, 7, 10, 9}
	// fmt.Println("Initial array: ", v1)
	// customSort.QuickSort(v1)
	// fmt.Println("Final result ", v1)
	// fmt.Printf("arraySorted = %v\n", customSort.IsSorted(v1))
	//
	// s := stack.LinkedStack{}
	// s.Push(4)
	// s.Push(5)
	// s.Push(10)
	// s.Push(5)
	// s.Push(20)
	// s.Push(50)
	// stack.Print(&s)
	//
	// fmt.Println(s.Min())

	h := binaryHeap.NewHeap(10)
	for i := 1; i <= 10; i++ {
		h.Insert(i)
	}
	h.PrintArray()
	fmt.Println(h.DeleteMax())
	h.PrintArray()

	// h.Sort()
	// h.PrintArray()
}
