package heap

import "container/heap"

type myHeap []any

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].(int) < (*h)[j].(int)
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *myHeap) Top() any {
	return (*h)[0]
}

func (h *myHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func topKHeap(nums []int, k int) *myHeap {
	h := &myHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Top().(int) {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return h
}
