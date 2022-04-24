package dash

type stack struct {
	values []string
	index  int
}

func newStackReverse(values []string) *stack {
	s := &stack{
		values: values,
		index:  len(values),
	}
	return s
}

func (s *stack) empty() bool {
	return s.index < len(s.values)
}

func (s *stack) pop() string {
	if s.index <= 0 {
		panic("pop from empty stack")
	}
	s.index--
	value := s.values[s.index]
	return value
}

func (s *stack) push(value string) {
	if s.index >= len(s.values) {
		panic("push to full stack")
	}
	s.values[s.index] = value
	s.index++
}
