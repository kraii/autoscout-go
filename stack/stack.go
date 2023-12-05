package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func EmptyStack[T any]() *Stack[T] {
	items := make([]T, 0, 10)
	return &Stack[T]{items}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) <= 0
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var nothing T
		return nothing, fmt.Errorf("pop on empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) Print() {
	print("[")
	for _, item := range s.items {
		print(item)
		print(",")
	}
	print("]")
	println()
}
