package src

func graphBFS(g *graphAdjList, startVet Vertex) []Vertex {
	res := make([]Vertex, 0)
	visited := make(map[Vertex]struct{})
	visited[startVet] = struct{}{}
	queue := make([]Vertex, 0)
	queue = append(queue, startVet)
	for len(queue) > 0 {
		vet := queue[0]
		queue = queue[1:]
		res = append(res, vet)
		for _, adjVet := range g.adjList[vet] {
			_, isExist := visited[adjVet]
			if !isExist {
				queue = append(queue, adjVet)
				visited[adjVet] = struct{}{}
			}
		}
	}
	return res
}
