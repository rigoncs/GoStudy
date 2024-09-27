package src

type arrayDeque struct {
	nums        []int
	queCapacity int
	queSize     int
	front       int
}

func newArrayDeque(capacity int) *arrayDeque {
	return &arrayDeque{
		nums:        make([]int, capacity),
		queCapacity: capacity,
		queSize:     0,
		front:       0,
	}
}

func (q *arrayDeque) size() int {
	return q.queSize
}

func (q *arrayDeque) isEmpty() bool {
	return q.queSize == 0
}

/*计算环形数组索引*/
func (q *arrayDeque) index(i int) int {
	return (i + q.queCapacity) % q.queCapacity
}

func (q *arrayDeque) pushFirst(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	q.front = q.index(q.front - 1)
	q.nums[q.front] = num
	q.queSize++
}

func (q *arrayDeque) pushLast(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := q.index(q.front + q.queSize)
	q.nums[rear] = num
	q.queSize++
}

func (q *arrayDeque) popFirst() any {
	e := q.peekFirst()
	if e == nil {
		return nil
	}
	q.front = q.index(q.front + 1)
	q.queSize--
	return e
}

func (q *arrayDeque) popLast() any {
	e := q.peekLast()
	if e == nil {
		return nil
	}
	q.queSize--
	return e
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
	last := q.index(q.front + q.queSize - 1)
	return q.nums[last]
}

func (q *arrayDeque) toSlice() []int {
	res := make([]int, q.queSize)
	for i, j := 0, q.front; i < q.queSize; i++ {
		res[i] = q.nums[q.index(j)]
		j++
	}
	return res
}
