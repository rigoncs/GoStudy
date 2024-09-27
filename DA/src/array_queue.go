package src

type arrayQueue struct {
	nums        []int
	queCapacity int
	queSize     int
	front       int
}

func newArrayQueue(capacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, capacity),
		queCapacity: capacity,
		queSize:     0,
		front:       0,
	}
}

func (q *arrayQueue) size() int {
	return q.queSize
}

func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayQueue) push(value int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := (q.front + q.queSize) % q.queCapacity
	q.nums[rear] = value
	q.queSize++
}

func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayQueue) pop() any {
	e := q.peek()
	if e == nil {
		return nil
	}
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return e
}

func (q *arrayQueue) toSlice() []int {
	rear := q.front + q.queSize
	if rear > q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}
