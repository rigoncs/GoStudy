package graph

type graphAdjMat struct {
	vertices []int
	adjMat   [][]int
}

func newGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}
	for _, edge := range edges {
		g.addEdge(edge[0], edge[1])
	}
	return g
}

func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

func (g *graphAdjMat) addVertex(val int) {
	n := g.size()
	g.vertices = append(g.vertices, val)
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

func (g *graphAdjMat) removeVertex(index int) {
	if index < 0 || index >= g.size() {
		return
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

func (g *graphAdjMat) addEdge(i, j int) {
	if i < 0 || i >= g.size() || j < 0 || j > +g.size() || i == j {
		return
	}
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

func (g *graphAdjMat) removeEdge(i, j int) {
	if i < 0 || i >= g.size() || j < 0 || j > +g.size() || i == j {
		return
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}
