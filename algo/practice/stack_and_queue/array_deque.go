package stackandqueue

type arrayDeque struct {
	nums        []int
	front       int
	queSize     int
	queCapacity int
}

func newArrayDeque(queCapacity int) *arrayDeque {
	return &arrayDeque{
		nums:        make([]int, queCapacity),
		front:       0,
		queSize:     0,
		queCapacity: queCapacity,
	}
}

func (q *arrayDeque) size() int {
	return q.queSize
}

func (q *arrayDeque) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayDeque) pushFirst(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	idx := (q.front + q.queCapacity - 1) % q.queCapacity
	q.nums[idx] = num
	q.front = idx
	q.queSize++
}

func (q *arrayDeque) pushLast(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	idx := (q.front + q.queSize) % q.queCapacity
	q.nums[idx] = num
	q.queSize++
}

func (q *arrayDeque) peekFirst() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayDeque) peekLast() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[(q.front+q.queSize-1)%q.queCapacity]
}

func (q *arrayDeque) popFirst() any {
	num := q.peekFirst()
	if num == nil {
		return nil
	}
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

func (q *arrayDeque) popLast() any {
	num := q.peekFirst()
	if num == nil {
		return nil
	}
	q.queSize--
	return num
}

func (q *arrayDeque) toSlice() []int {
	res := make([]int, q.queSize)
	for i, j := 0, q.front; i < q.queSize; i++ {
		res[i] = q.nums[j%q.queCapacity]
		j++
	}
	return res
}
