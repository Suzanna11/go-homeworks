package main

import (
	"fmt"
)

type graph struct {
	NumOfVertices int
	Edges         [][]int
}
func main(){
	var n int
	fmt.Scanln(&n,newGraph(n))

}
func newGraph(n int) *graph {
	return &graph{
		NumOfVertices: n,
		Edges:         make([][]int, n),
	}
}
func (g *graph) addEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
}

