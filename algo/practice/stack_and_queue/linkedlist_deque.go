package stackandqueue

import "container/list"

type linkedListDeque struct {
	data *list.List
}

func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}

func (dq *linkedListDeque) size() int {
	return dq.data.Len()
}

func (dq *linkedListDeque) isEmpty() bool {
	return dq.data.Len() == 0
}

func (dq *linkedListDeque) pushFirst(value int) {
	dq.data.PushFront(value)
}

func (dq *linkedListDeque) pushLast(value int) {
	dq.data.PushBack(value)
}

func (dq *linkedListDeque) peekFirst() any {
	if dq.isEmpty() {
		return nil
	}
	e := dq.data.Front()
	return e.Value
}

func (dq *linkedListDeque) peekLast() any {
	if dq.isEmpty() {
		return nil
	}
	e := dq.data.Back()
	return e.Value
}

func (dq *linkedListDeque) popFirst() any {
	if dq.isEmpty() {
		return nil
	}
	e := dq.data.Front()
	dq.data.Remove(e)
	return e.Value
}

func (dq *linkedListDeque) popLast() any {
	if dq.isEmpty() {
		return nil
	}
	e := dq.data.Back()
	dq.data.Remove(e)
	return e.Value
}

func (dq *linkedListDeque) toList() *list.List {
	return dq.data
}
