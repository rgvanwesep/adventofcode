package day1

import (
	"regexp"
	"slices"
	"strconv"
)

func Dist(x, y uint64) uint64 {
	if x < y {
		return y - x
	}
	return x - y
}

func ParseLocations(inputs []string) [2][]uint64 {
	locations := [2][]uint64{make([]uint64, 0), make([]uint64, 0)}
	pattern := regexp.MustCompile(`(\d+)\s+(\d+)`)
	for _, input := range inputs {
		match := pattern.FindStringSubmatch(input)
		for i := range locations {
			if num, err := strconv.ParseUint(match[i+1], 10, 64); err == nil {
				locations[i] = append(locations[i], num)
			}
		}
	}
	return locations
}

func SumDistances(inputs []string) uint64 {
	var sum uint64 = 0
	locations := ParseLocations(inputs)
	for i := range locations {
		slices.Sort(locations[i])
	}
	for i := 0; i < len(locations[0]); i++ {
		sum += Dist(locations[0][i], locations[1][i])
	}
	return sum
}

func Count(x []uint64) map[uint64]uint64 {
	counts := make(map[uint64]uint64)
	for _, i := range x {
		counts[i]++
	}
	return counts
}

func CalcSimilarity(inputs []string) uint64 {
	var similarity uint64 = 0
	locations := ParseLocations(inputs)
	counts := Count(locations[1])
	for _, i := range locations[0] {
		similarity += i * counts[i]
	}
	return similarity
}
