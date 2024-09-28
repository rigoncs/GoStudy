package arrayandlinkedlist

type myList struct {
	arr         []int
	arrSize     int
	arrCapacity int
	extendRatio int
}

func newMyList() *myList {
	return &myList{
		arr:         make([]int, 10),
		arrSize:     0,
		arrCapacity: 10,
		extendRatio: 2,
	}
}

func (l *myList) size() int {
	return l.arrSize
}

func (l *myList) capacity() int {
	return l.arrCapacity
}

func (l *myList) get(idx int) int {
	if idx < 0 || idx >= l.size() {
		return -1
	}
	return l.arr[idx]
}

func (l *myList) set(idx, num int) {
	if idx < 0 || idx >= l.size() {
		panic("Index out of bounds!")
	}
	l.arr[idx] = num
}

func (l *myList) add(num int) {
	if l.arrSize == l.arrCapacity {
		l.extend()
	}
	l.arr[l.arrSize] = num
	l.arrSize++
}

func (l *myList) insert(num, idx int) {
	if idx < 0 || idx >= l.size() {
		panic("Index out of bounds!")
	}
	if l.arrSize == l.arrCapacity {
		l.extend()
	}
	for i := l.size(); i > idx; i-- {
		l.arr[i] = l.arr[i-1]
	}
	l.arr[idx] = num
	l.arrSize++
}

func (l *myList) extend() {
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	l.arrCapacity = len(l.arr)
}

func (l *myList) remove(index int) int {
	if index < 0 || index >= l.size() {
		panic("Index out of bounds!")
	}
	res := l.arr[index]
	for i := index; i < l.size()-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.arrSize--
	return res
}

func (l *myList) toArray() []int {
	return l.arr[:l.arrSize]
}
