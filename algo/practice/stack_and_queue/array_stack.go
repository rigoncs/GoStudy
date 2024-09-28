package stackandqueue

type arrayStack struct {
	data []int
}

func newArrayStack() *arrayStack {
	return &arrayStack{data: make([]int, 0, 16)}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *arrayStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *arrayStack) push(num int) {
	s.data = append(s.data, num)
}

func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *arrayStack) toSlice() []int {
	return s.data
}
