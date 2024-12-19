package day19

import "strings"

type basePatterns struct {
	patterns []string
}

func (b basePatterns) isPossible(pattern string) bool {
	possibles := make([]bool, len(pattern)+1)
	possibles[0] = true
	for i := 1; i < len(possibles); i++ {
		for _, basePattern := range b.patterns {
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
				}
			}
		}
	}
	return possibles[len(pattern)]
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
