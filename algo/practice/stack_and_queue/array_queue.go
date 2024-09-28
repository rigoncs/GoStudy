package stackandqueue

type arrayQueue struct {
	data        []int
	queSize     int
	front       int
	queCapacity int
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		data:        make([]int, queCapacity),
		queSize:     0,
		queCapacity: queCapacity,
		front:       0,
	}
}

func (q *arrayQueue) size() int {
	return q.queSize
}

func (q *arrayQueue) isEmpty() bool {
	return q.size() == 0
}

func (q *arrayQueue) push(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := (q.front + q.queSize) % q.queCapacity
	q.data[rear] = num
	q.queSize++
}

func (q *arrayQueue) peek() any {
	if q.queSize == 0 {
		return nil
	}
	return q.data[q.front]
}

func (q *arrayQueue) pop() any {
	if q.queSize == 0 {
		return nil
	}
	num := q.data[q.front]
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

func (q *arrayQueue) toSlice() []int {
	rear := q.front + q.queSize
	if rear > q.queCapacity {
		rear %= q.queCapacity
		return append(q.data[q.front:], q.data[:rear]...)
	}
	return q.data[q.front:rear]
}
