package src

import "sort"

type Item struct {
	v int
	w int
}

func fractionalKnapsack(weg, val []int, cap int) float64 {
	items := make([]Item, len(weg))
	for i := 0; i < len(weg); i++ {
		items[i] = Item{val[i], weg[i]}
	}
	sort.Slice(items, func(i, j int) bool {
		return float64(items[i].v)/float64(items[i].w) > float64(items[j].v)/float64(items[j].w)
	})
	res := 0.0
	for _, item := range items {
		if item.w <= cap {
			res += float64(item.v)
			cap -= item.w
		} else {
			res += float64(item.v) / float64(item.w) * float64(cap)
			break
		}
	}
	return res
}
