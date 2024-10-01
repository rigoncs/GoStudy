package graph

import (
	. "practice/pkg"
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
	g.adjList[vet1] = DeleteSliceElms(g.adjList[vet1], vet2)
	g.adjList[vet2] = DeleteSliceElms(g.adjList[vet2], vet1)
}

func (g *graphAdjList) addVertex(vet Vertex) {
	if _, ok := g.adjList[vet]; ok {
		return
	}
	g.adjList[vet] = make([]Vertex, 0)
}

func (g *graphAdjList) removeVertex(vet Vertex) {
	if _, ok := g.adjList[vet]; !ok {
		panic("Error")
	}
	delete(g.adjList, vet)
	for v, list := range g.adjList {
		g.adjList[v] = DeleteSliceElms(list, vet)
	}
}
