package graph

import . "practice/pkg"

func graphBFS(g *graphAdjList, startVet Vertex) []Vertex {
	visited := make(map[Vertex]struct{})
	res := make([]Vertex, 0)
	visited[startVet] = struct{}{}
	queue := make([]Vertex, 0)
	queue = append(queue, startVet)
	for len(queue) > 0 {
		vet := queue[0]
		queue = queue[1:]
		res = append(res, vet)
		for _, adjVet := range g.adjList[vet] {
			if _, isVisited := visited[adjVet]; !isVisited {
				queue = append(queue, adjVet)
				visited[adjVet] = struct{}{}
			}
		}
	}
	return res
}
