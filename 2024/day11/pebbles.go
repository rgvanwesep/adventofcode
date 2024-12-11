package day11

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func Split(x int) []int {
	digits := []int{}
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	slices.Reverse(digits)
	return digits
}

func Join(digits []int) int {
	nDigits := len(digits)
	x := digits[0]
	for i := 1; i < nDigits; i++ {
		x *= 10
		x += digits[i]
	}
	return x
}

func CountAncestors(pebble int, nBlinks int) int {
	if nBlinks == 0 {
		return 1
	}
	digits := Split(pebble)
	nDigits := len(digits)
	newPebbles := make([]int, 0, 2)
	switch {
	case pebble == 0:
		newPebbles = append(newPebbles, 1)
	case nDigits%2 == 0:
		newPebbles = append(newPebbles, Join(digits[:nDigits/2]), Join(digits[nDigits/2:]))
	default:
		newPebbles = append(newPebbles, pebble*2024)
	}
	count := 0
	for _, p := range newPebbles {
		count += CountAncestors(p, nBlinks - 1)
	}
	return count
}

func CountPebbles(inputs []string) int {
	const nBlinks = 25
	input := inputs[0]
	split := strings.Split(input, " ")
	count := 0
	for _, s := range split {
		if pebble, err := strconv.Atoi(s); err == nil {
			count += CountAncestors(pebble, nBlinks)
		} else {
			log.Panicf("Could not parse %q", input)
		}
	}
	return count
}
