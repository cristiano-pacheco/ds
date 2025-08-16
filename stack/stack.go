package stack

type Stack[T any] struct {
	items []T
	top   int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
		top:   0,
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.top = s.top + 1
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[s.top]
	s.top = s.top - 1
	return item
}

func (s *Stack[T]) Top() T {
	return s.Peak()
}

func (s *Stack[T]) Peak() T {
	return s.items[s.top]
}
