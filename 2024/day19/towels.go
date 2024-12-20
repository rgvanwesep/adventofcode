package day19

import "strings"

type basePatterns struct {
	patterns []string
}

func (b basePatterns) countPossibles(pattern string) int {
	possibles := make([]bool, len(pattern)+1)
	possibles[0] = true
	matches := make([][]int, len(pattern)+1)
	counts := make([]int, len(pattern)+1)
	counts[0] = 1
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
		for _, k := range matches[i] {
			counts[i] += counts[i-len(b.patterns[k])]
		}
	}
	return counts[len(pattern)]
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
