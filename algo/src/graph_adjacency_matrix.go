package src

import "fmt"

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
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
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
	if index > g.size() {
		return
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

func (g *graphAdjMat) addEdge(i, j int) {
	if i < 0 || j < 0 || i > g.size() || j > g.size() || i == j {
		fmt.Errorf("%s", "Index Out of Bounds Exception")
	}
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

func (g *graphAdjMat) removeEdge(i, j int) {
	if i < 0 || j < 0 || i > g.size() || j > g.size() || i == j {
		fmt.Errorf("%s", "Index Out of Bounds Exception")
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

func (g *graphAdjMat) print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t邻接矩阵 = \n")
	for i := range g.adjMat {
		fmt.Printf("\t\t\t%v\n", g.adjMat[i])
	}
}
