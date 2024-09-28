package src

type Vertex struct {
	Val int
}

func newVertex(val int) Vertex {
	return Vertex{
		Val: val,
	}
}

func valsToVets(nums []int) []Vertex {
	vets := make([]Vertex, len(nums))
	for i := 0; i < len(vets); i++ {
		vets[i] = newVertex(nums[i])
	}
	return vets
}

func vetsToVals(vets []Vertex) []int {
	vals := make([]int, len(vets))
	for i := 0; i < len(vals); i++ {
		vals[i] = vets[i].Val
	}
	return vals
}

func deleteSlicedElms[T any](a []T, elms ...T) []T {
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
