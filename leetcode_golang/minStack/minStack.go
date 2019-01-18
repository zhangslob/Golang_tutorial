package minStack

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	min := x
	if len(s.stack) > 0 && s.GetMin() < x {
		min = s.GetMin()
	}
	s.stack = append(s.stack, item{min: min, x: x})
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].min
}

func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}



