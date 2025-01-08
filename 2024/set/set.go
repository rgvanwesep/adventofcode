package set

import "iter"

type Set[T comparable] interface {
	Add(T)
	All() iter.Seq[T]
	Clear()
	Clone() Set[T]
	Contains(T) bool
	Len() int
	Remove(T)
}

func NewSet[T comparable]() Set[T] {
	s := new(setMap[T])
	s.values = map[T]struct{}{}
	return s
}

type setMap[T comparable] struct {
	values map[T]struct{}
}

func (s *setMap[T]) Add(value T) {
	s.values[value] = struct{}{}
}

func (s *setMap[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range s.values {
			if !yield(value) {
				break
			}
		}
	}
}

func (s *setMap[T]) Clear() {
	s.values = map[T]struct{}{}
}

func (s *setMap[T]) Clone() Set[T] {
	sClone := new(setMap[T])
	sClone.values = map[T]struct{}{}
	for value := range s.values {
		sClone.values[value] = struct{}{}
	}
	return sClone
}

func (s *setMap[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}

func (s *setMap[T]) Len() int {
	return len(s.values)
}

func (s *setMap[T]) Remove(value T) {
	delete(s.values, value)
}
