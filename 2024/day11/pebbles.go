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

func CountPebbles(inputs []string) int {
	const nBlinks = 25
	input := inputs[0]
	split := strings.Split(input, " ")
	pebbles := make([]int, len(split))
	for i, s := range split {
		if pebble, err := strconv.Atoi(s); err == nil {
			pebbles[i] = pebble
		} else {
			log.Panicf("Could not parse %q", input)
		}
	}
	for range nBlinks {
		newPebbles := []int{}
		for _, pebble := range pebbles {
			digits := Split(pebble)
			nDigits := len(digits)
			switch {
			case pebble == 0:
				newPebbles = append(newPebbles, 1)
			case nDigits%2 == 0:
				newPebbles = append(newPebbles, Join(digits[:nDigits/2]), Join(digits[nDigits/2:]))
			default:
				newPebbles = append(newPebbles, pebble*2024)
			}
		}
		pebbles = newPebbles
	}
	return len(pebbles)
}
