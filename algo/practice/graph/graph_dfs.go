package graph

import (
	. "practice/pkg"
	"strconv"
)

func dfs(g *graphAdjList, visited map[Vertex]struct{}, res *[]Vertex, vet Vertex) {
	*res = append(*res, vet)
	visited[vet] = struct{}{}
	for _, adjVet := g.adjList[vet] {
		if _, isVisited := visited[adjVet]; !isVisited {
			dfs(g, visited, res, adjVet)
		}
	}
}

func graphDFS(g *graphAdjList, startVet Vertex) []Vertex {
	res := make([]Vertex, 0)
	visited := make(map[Vertex]struct{})
	dfs(g, visited, &res, startVet)
	return res
}
