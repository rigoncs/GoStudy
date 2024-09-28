package src

import (
	"fmt"
	"strconv"
	"strings"
)

type graphAdjList struct {
	adjList map[Vertex][]Vertex
}

func newGraphAdjList(edges [][]Vertex) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

func (g *graphAdjList) size() int {
	return len(g.adjList)
}

func (g *graphAdjList) addEdge(vet1, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("Error")
	}
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

func (g *graphAdjList) removeEdge(vet1, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("Error")
	}
	g.adjList[vet1] = deleteSlicedElms(g.adjList[vet1], vet2)
	g.adjList[vet2] = deleteSlicedElms(g.adjList[vet2], vet1)
}

func (g *graphAdjList) addVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	g.adjList[vet] = make([]Vertex, 0)
}

func (g *graphAdjList) removeVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if !ok {
		panic("Error")
	}
	delete(g.adjList, vet)
	for v, list := range g.adjList {
		g.adjList[v] = deleteSlicedElms(list, vet)
	}
}

func (g *graphAdjList) print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for _, vet := range v {
			builder.WriteString(strconv.Itoa(vet.Val) + " ")
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}
