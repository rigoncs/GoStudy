package src

func dfs(g *graphAdjList, visited map[Vertex]struct{}, res *[]Vertex, vet Vertex) {
	*res = append(*res, vet)
	visited[vet] = struct{}{}
	for _, adjVet := range g.adjList[vet] {
		_, isExist := visited[adjVet]
		if !isExist {
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
