package src

import "container/list"

type linkedListQueue struct {
	data *list.List
}

func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

func (q *linkedListQueue) push(value any) {
	q.data.PushBack(value)
}

func (q *linkedListQueue) pop() any {
	if q.isEmpty() {
		return nil
	}
	e := q.data.Front()
	q.data.Remove(e)
	return e.Value
}

func (q *linkedListQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	e := q.data.Front()
	return e.Value
}

func (q *linkedListQueue) size() int {
	return q.data.Len()
}

func (q *linkedListQueue) isEmpty() bool {
	return q.data.Len() == 0
}

func (q *linkedListQueue) toList() *list.List {
	return q.data
}
