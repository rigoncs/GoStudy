package pkg

type Vertex struct {
	Val int
}

func NewVertex(val int) Vertex {
	return Vertex{Val: val}
}

func ValsToVets(vals []int) []Vertex {
	res := make([]Vertex, len(vals))
	for i, val := range vals {
		res[i] = NewVertex(val)
	}
	return res
}

func VetsToVals(vets []Vertex) []int {
	res := make([]int, len(vets))
	for i, vet := range vets {
		res[i] = vet.Val
	}
	return res
}

func DeleteSliceElms[T any](a []T, elms ...T) []T {
	if len(a) == 0 || len(elms) == 0 {
		return a
	}
	m := make(map[any]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	res := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
