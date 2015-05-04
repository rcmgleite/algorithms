package graph

import (
	"container/list"
	"fmt"

	"github.com/rcmgleite/algorithms/queue"
	"github.com/rcmgleite/algorithms/stack"
)

// Graph representation
// TODO - implement custom list (study)
type Graph struct {
	vertexCount int
	adjList     []*list.List
}

// NewGraph is the constructor for graph
func NewGraph(vertexCount int) *Graph {
	adjList := make([]*list.List, vertexCount) // construct slice
	for i := 0; i < vertexCount; i++ {
		adjList[i] = list.New() // create each list of the slice
	}

	return &Graph{vertexCount: vertexCount, adjList: adjList}
}

// AddEdge to the graph
func (g *Graph) AddEdge(v int, w int) {
	g.adjList[v].PushBack(w)
	g.adjList[w].PushBack(v)
}

// AdjIterator will return an iterator for the adjList of the vertex v
func (g *Graph) AdjIterator(v int) *list.Element {
	return g.adjList[v].Front()
}

// PrintAdjList ...
func (g *Graph) PrintAdjList(v int) {
	for iterator := g.AdjIterator(v); iterator != nil; iterator = iterator.Next() {
		fmt.Printf(" %v ", iterator.Value)
	}
	fmt.Println("")
}

// DfsPaths depth-first search - used on maze exploration
// Maze:
//  1) Vertex of the graph = intersection
//  2) Edge of the graph = passage
// Goal: Systematically search through a graph
// Ideia: Mimic maze exploration
type DfsPaths struct {
	Marked []bool
	EdgeTo []int
	s      int //initial vertex (source)
	g      Graph
}

// NewDfsPaths constructor. It executes the Apply method that runs dfs algorithm
func NewDfsPaths(g Graph, s int) *DfsPaths {
	edgeTo := make([]int, g.vertexCount)
	for i := 0; i < g.vertexCount; i++ {
		edgeTo[i] = -1
	}
	dfs := &DfsPaths{Marked: make([]bool, g.vertexCount), EdgeTo: edgeTo, s: s, g: g}
	dfs.apply(s)
	return dfs
}

// apply ... v = destination
func (dfs *DfsPaths) apply(v int) {
	dfs.Marked[v] = true
	for w := dfs.g.AdjIterator(v); w != nil; w = w.Next() {
		if !dfs.Marked[w.Value.(int)] {
			dfs.apply(w.Value.(int)) // depth-first search recusrive calls act like a stack
			dfs.EdgeTo[w.Value.(int)] = v
		}
	}
}

// HasPathTo returns true if there is a path connecting the source to v
func (dfs *DfsPaths) HasPathTo(v int) bool {
	return dfs.Marked[v]
}

// PathTo return the path from s to v
func (dfs *DfsPaths) PathTo(v int) *stack.LinkedStack {
	if !dfs.HasPathTo(v) {
		return nil
	}

	path := &stack.LinkedStack{}
	for i := v; i != dfs.s; i = dfs.EdgeTo[i] {
		path.Push(i)
	}
	path.Push(dfs.s)
	return path
}

// BfsPaths ... same as DfsPaths
type BfsPaths struct {
	Marked []bool
	EdgeTo []int
	DistTo []int
	Queue  queue.LinkedQueue
	s      int //initial vertex (source)
	g      Graph
}

// NewBfsPaths constructor. It executes the Apply method that runs dfs algorithm
func NewBfsPaths(g Graph, s int) *BfsPaths {
	edgeTo := make([]int, g.vertexCount)
	distTo := make([]int, g.vertexCount)
	for i := 0; i < g.vertexCount; i++ {
		edgeTo[i] = -1
		distTo[i] = -1
	}
	bfs := &BfsPaths{Marked: make([]bool, g.vertexCount), EdgeTo: edgeTo, DistTo: distTo, Queue: queue.LinkedQueue{}, s: s, g: g}
	bfs.apply(s)                         // Apply breadth-first search algorithm
	for i := 0; i < g.vertexCount; i++ { // calculate DistTo from s to every other vertex
		if i == s {
			continue
		}
		path := bfs.ShortestPathTo(i)
		bfs.DistTo[i] = path.Size() - 1 // -1 to remove the vertex s from counting
	}

	return bfs
}

// Apply breadth-first search - usade on mazes too
// facebook hacker cup laser maze - shortest path algorithm
// It utilizes a QUEUE instead of a stack (DFS)
func (bfs *BfsPaths) apply(v int) {
	bfs.Marked[v] = true
	bfs.Queue.Enqueue(v) //a

	for !bfs.Queue.IsEmpty() { // until queue is empty
		v = bfs.Queue.Dequeue()
		for w := bfs.g.AdjIterator(v); w != nil; w = w.Next() { //search through every adj node
			if !bfs.Marked[w.Value.(int)] {
				bfs.Queue.Enqueue(w.Value.(int))
				bfs.EdgeTo[w.Value.(int)] = v
				bfs.Marked[w.Value.(int)] = true
			}
		}
	}
}

// HasPathTo ...
func (bfs *BfsPaths) HasPathTo(v int) bool {
	return bfs.Marked[v]
}

// ShortestPathTo ...
func (bfs *BfsPaths) ShortestPathTo(v int) *stack.LinkedStack {
	if !bfs.HasPathTo(v) {
		return nil
	}

	path := &stack.LinkedStack{}
	for i := v; i != bfs.s; i = bfs.EdgeTo[i] {
		path.Push(i)
	}
	path.Push(bfs.s)
	return path
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////// Connected Components ////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
// Using DFS we want to get all the connected componenets of a graph
//   -> Solution: run DFS for all the vertex and find the CC

// ConnectedComponents ...
type ConnectedComponents struct {
	Marked []bool
	Ids    []int
	Count  int
	G      *Graph
}

// NewCC ...
func NewCC(g *Graph) *ConnectedComponents {
	marked := make([]bool, g.vertexCount)
	ids := make([]int, g.vertexCount)
	cc := &ConnectedComponents{Marked: marked, Ids: ids, Count: 0, G: g}

	for v := 0; v < g.vertexCount; v++ {
		if !marked[v] {
			cc.dfs(v)
			cc.Count++
		}
	}

	return cc
}

func (cc *ConnectedComponents) dfs(v int) {
	cc.Marked[v] = true
	cc.Ids[v] = cc.Count
	for w := cc.G.adjList[v].Front(); w != nil; w = w.Next() {
		if !cc.Marked[w.Value.(int)] {
			cc.dfs(w.Value.(int))
		}
	}
}

// IsConnected ... after building cc object, verify if 2 vertex are connected
// in constant time
func (cc *ConnectedComponents) IsConnected(v, w int) bool {
	return cc.Ids[v] == cc.Ids[w]
}
