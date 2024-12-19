package day19

import "strings"

type stack[T any] struct {
	values []T
}

func newStack[T any]() stack[T] {
	return stack[T]{
		values: []T{},
	}
}

func (s *stack[T]) push(v T) {
	s.values = append(s.values, v)
}

func (s *stack[T]) pop() (v T, ok bool) {
	size := len(s.values)
	if size > 0 {
		v = s.values[size-1]
		s.values = s.values[:size-1]
		ok = true
	}
	return
}

func (s *stack[T]) clear() {
	s.values = []T{}
}

type combination struct {
	currIndex int
	values    []int
}

type basePatterns struct {
	patterns []string
}

func (b basePatterns) allPossibles(pattern string) [][]int {
	possibles := make([]bool, len(pattern)+1)
	possibles[0] = true
	matches := make([][]int, len(pattern)+1)
	for i := 1; i < len(possibles); i++ {
		for k, basePattern := range b.patterns {
			if i-len(basePattern) >= 0 && possibles[i-len(basePattern)] {
				baseMatch := true
				for j := 0; j < len(basePattern); j++ {
					if basePattern[j] != pattern[i-len(basePattern)+j] {
						baseMatch = false
						break
					}
				}
				if baseMatch {
					possibles[i] = true
					matches[i] = append(matches[i], k)
				}
			}
		}
	}
	combinations := [][]int{}
	s := newStack[combination]()
	for _, k := range matches[len(pattern)] {
		s.push(combination{len(pattern), []int{k}})
	}
	for {
		if c, ok := s.pop(); ok {
			prevIndex := c.currIndex - len(b.patterns[c.values[0]])
			if prevIndex > 0 {
				for _, k := range matches[prevIndex] {
					s.push(combination{prevIndex, append([]int{k}, c.values...)})
				}
			} else {
				combinations = append(combinations, c.values)
			}
		} else {
			break
		}
	}
	return combinations
}

func (b basePatterns) isPossible(pattern string) bool {
	combinations := b.allPossibles(pattern)
	return len(combinations) > 0
}

func parsePatterns(inputs []string) (basePatterns, []string) {
	return basePatterns{strings.Split(inputs[0], ", ")}, inputs[2:]
}

func CountPossible(inputs []string) int {
	count := 0
	base, patterns := parsePatterns(inputs)
	for _, pattern := range patterns {
		if base.isPossible(pattern) {
			count++
		}
	}
	return count
}

func SumCombinations(inputs []string) int {
	sum := 0
	base, patterns := parsePatterns(inputs)
	for _, pattern := range patterns {
		sum += len(base.allPossibles(pattern))
	}
	return sum
}
