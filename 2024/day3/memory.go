package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func SumMul(inputs []string) uint64 {
	var (
		sum          uint64 = 0
		err          error
		factor, term uint64
	)
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, input := range inputs {
		matches := pattern.FindAllStringSubmatch(input, -1)
		for _, groups := range matches {
			term = 1
			for i := 1; i <= 2; i++ {
				if factor, err = strconv.ParseUint(groups[i], 10, 64); err == nil {
					term *= factor
				} else {
					panic(fmt.Sprintf("Unparsable factor in match %q", groups[0]))
				}
			}
			sum += term
		}
	}
	return sum
}

func SumConditionalMul(inputs []string) uint64 {
	var (
		sum          uint64 = 0
		enabled      bool   = true
		err          error
		factor, term uint64
	)
	pattern := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	for _, input := range inputs {
		matches := pattern.FindAllStringSubmatch(input, -1)
		for _, groups := range matches {
			if groups[0] == "do()" {
				enabled = true
			} else if groups[0] == "don't()" {
				enabled = false
			} else if enabled {
				term = 1
				for i := 2; i <= 3; i++ {
					if factor, err = strconv.ParseUint(groups[i], 10, 64); err == nil {
						term *= factor
					} else {
						panic(fmt.Sprintf("Unparsable factor in match %q", groups[0]))
					}
				}
				sum += term
			}
		}
	}
	return sum
}
