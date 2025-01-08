package set

import (
	"fmt"
	"iter"
	"strings"
)

type Set[T comparable] interface {
	Add(T)
	All() iter.Seq[T]
	Clear()
	Clone() Set[T]
	Contains(T) bool
	Len() int
	Remove(T)
	String() string
}

func NewSet[T comparable]() Set[T] {
	s := new(setMap[T])
	s.values = map[T]struct{}{}
	return s
}

func Equals[T comparable](x, y Set[T]) bool {
	if x.Len() == y.Len() {
		for v := range x.All() {
			if !y.Contains(v) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func Intersection[T comparable](x, y Set[T]) Set[T] {
	var left, right Set[T]
	z := NewSet[T]()
	if x.Len() < y.Len() {
		left = x
		right = y
	} else {
		left = y
		right = x
	}
	for v := range left.All() {
		if right.Contains(v) {
			z.Add(v)
		}
	}
	return z
}

func Union[T comparable](x, y Set[T]) Set[T] {
	z := NewSet[T]()
	for v := range x.All() {
		z.Add(v)
	}
	for v := range y.All() {
		z.Add(v)
	}
	return z
}

func Difference[T comparable](x, y Set[T]) Set[T] {
	z := x.Clone()
	for v := range y.All() {
		z.Remove(v)
	}
	return z
}

func SymmetricDifference[T comparable](x, y Set[T]) Set[T] {
	z := NewSet[T]()
	for v := range x.All() {
		if !y.Contains(v) {
			z.Add(v)
		}
	}
	for v := range y.All() {
		if !x.Contains(v) {
			z.Add(v)
		}
	}
	return z
}

func CartesianProduct[T comparable](x, y Set[T]) Set[[2]T] {
	z := NewSet[[2]T]()
	for vx := range x.All() {
		for vy := range y.All() {
			z.Add([2]T{vx, vy})
		}
	}
	return z
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

func (s setMap[T]) String() string {
	parts := make([]string, s.Len())
	i := 0
	for v := range s.All() {
		parts[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings.Join([]string{"{", strings.Join(parts, ", "), "}"}, "")
}
