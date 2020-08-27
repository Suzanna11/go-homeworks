package main

import (
	"fmt"
)

type node struct
{
    vertex int
    next *node
}

type graph struct
{
    numVertices int
    list *adj
}

func (*graph) graph(vertex int){
	v := vertex
	var list = new([v]int)
	var adj = list


}

