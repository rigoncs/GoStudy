package heap

type minHeap struct {
	data []any
}

func newMinHeap(nums []any) *minHeap {
	h := &minHeap{data: nums}
	for i := h.parent(h.size() - 1); i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *minHeap) left(i int) int {
	return 2*i + 1
}

func (h *minHeap) right(i int) int {
	return 2*i + 2
}

func (h *minHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *minHeap) size() int {
	return len(h.data)
}

func (h *minHeap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *minHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *minHeap) peek() any {
	return h.data[0]
}

func (h *minHeap) push(val any) {
	h.data = append(h.data, val)
	h.siftUp(h.size() - 1)
}

func (h *minHeap) siftUp(i int) {
	for true {
		p := h.parent(i)
		if p < 0 || h.data[p].(int) < h.data[i].(int) {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *minHeap) pop() any {
	if h.isEmpty() {
		return nil
	}
	val := h.data[0]
	h.swap(0, h.size()-1)
	h.data = h.data[:h.size()-1]
	h.siftDown(0)
	return val
}

func (h *minHeap) siftDown(i int) {
	for true {
		l, r, min := h.left(i), h.right(i), i
		if l < h.size() && h.data[l].(int) < h.data[min].(int) {
			min = l
		}
		if r < h.size() && h.data[r].(int) < h.data[min].(int) {
			min = r
		}
		if min == i {
			break
		}
		h.swap(i, min)
		i = min
	}
}
