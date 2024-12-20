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

type basePatterns struct {
	patterns []string
}

func (b basePatterns) countPossibles(pattern string) int {
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
	matchesByLen := make([]map[int]int, len(matches))
	for i, m := range matches {
		matchesByLen[i] = map[int]int{}
		for _, k := range m {
			matchesByLen[i][len(b.patterns[k])]++
		}
	}
	count := 0
	s := newStack[[3]int]()
	for l, c := range matchesByLen[len(pattern)] {
		s.push([3]int{len(pattern), l, c})
	}
	for {
		if pair, ok := s.pop(); ok {
			prevIndex := pair[0] - pair[1]
			if prevIndex > 0 {
				for l, c := range matchesByLen[prevIndex] {
					s.push([3]int{prevIndex, l, pair[2] * c})
				}
			} else {
				count += pair[2]
			}
		} else {
			break
		}
	}
	return count
}

func (b basePatterns) isPossible(pattern string) bool {
	return b.countPossibles(pattern) > 0
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
		sum += base.countPossibles(pattern)
	}
	return sum
}
