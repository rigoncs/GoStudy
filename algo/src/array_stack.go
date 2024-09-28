package src

type arrayStack struct {
	data []int
}

func newArrayStack() *arrayStack {
	return &arrayStack{
		data: make([]int, 0, 16),
	}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return s.size() == 0
}

func (s *arrayStack) push(value int) {
	s.data = append(s.data, value)
}

func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data[len(s.data)-1]
	return e
}

func (s *arrayStack) pop() any {
	e := s.peek()
	s.data = s.data[:len(s.data)-1]
	return e
}

func (s *arrayStack) toSlice() []int {
	return s.data
}
