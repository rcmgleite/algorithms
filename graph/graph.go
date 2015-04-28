package graph

import (
	"container/list"
	"fmt"

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
			dfs.apply(w.Value.(int))
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

// Bfs breadth-first search - usade on mazes too
// facebook hacker cup laser maze
func Bfs() {
	//TODO
}
