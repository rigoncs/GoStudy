package src

type myList struct {
	arrCapacity int
	arrSize     int
	arr         []int
	extendRatio int
}

func newMyList() *myList {
	return &myList{
		arrCapacity: 10,
		arrSize:     0,
		arr:         make([]int, 10),
		extendRatio: 2,
	}
}

func (l *myList) size() int {
	return l.arrSize
}

func (l *myList) capacity() int {
	return l.arrCapacity
}

func (l *myList) get(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	return l.arr[index]
}

func (l *myList) set(index int, value int) {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	l.arr[index] = value
}

// 尾部添加元素
func (l *myList) add(num int) {
	if l.arrSize < l.arrCapacity {
		l.arr[l.arrSize] = num
		l.arrSize++
	} else {
		l.extendCapacity()
		l.arr[l.arrSize] = num
		l.arrSize++
	}
}

func (l *myList) extendCapacity() {
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	l.arrCapacity = len(l.arr)
}

func (l *myList) insert(index, num int) {
	if index < 0 || index >= l.arrSize {
		panic("索引错误")
	}
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	for j := l.arrSize; j > index; j-- {
		l.arr[j] = l.arr[j-1]
	}
	l.arr[index] = num
	l.arrSize++
}

func (l *myList) remove(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("索引错误")
	}
	num := l.arr[index]
	for j := index; j < l.arrSize-1; j++ {
		l.arr[j] = l.arr[j+1]
	}
	l.arrSize--
	return num
}

func (l *myList) toArray() []int {
	return l.arr[:l.arrSize]
}
