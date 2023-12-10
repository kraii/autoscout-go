package queue

import "fmt"

type Queue[T any] struct {
	items []T
}

func EmptyQueue[T any]() *Queue[T] {
	items := make([]T, 0, 10)
	return &Queue[T]{items}
}

func (s *Queue[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.items) <= 0
}

func (s *Queue[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var nothing T
		return nothing, fmt.Errorf("pop on empty queue")
	}
	item := s.items[0]
	s.items = s.items[1:]
	//item := s.items[len(s.items)-1]
	//s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Queue[T]) Print() {
	print("[")
	for _, item := range s.items {
		print(item)
		print(",")
	}
	print("]")
	println()
}
