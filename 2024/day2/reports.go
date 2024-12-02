package day2

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SAFE = iota
	UNSAFE
	INCREASING
	DECREASING
	STEADY
	UNSET
)

func CountSafeReports(inputs []string) uint64 {
	var (
		count         uint64 = 0
		level         int64
		diff          int64
		isDiffInRange bool
		direction     int
		err           error
	)

	reports := make([][]int64, 0)
	for _, input := range inputs {
		inputSplit := strings.Split(input, " ")
		levels := make([]int64, len(inputSplit))
		for i, s := range inputSplit {
			if level, err = strconv.ParseInt(s, 10, 64); err == nil {
				levels[i] = level
			} else {
				panic(fmt.Sprintf("Unparseable level in input: %q\n", input))
			}
		}
		reports = append(reports, levels)
	}

	ratings := make([]int, len(reports))
	for i, levels := range reports {
		direction = UNSET
		ratings[i] = UNSET
		for j := 0; j < len(levels)-1; j++ {
			diff = levels[j+1] - levels[j]
			if direction == UNSET {
				if diff < 0 {
					direction = DECREASING
				} else if diff > 0 {
					direction = INCREASING
				} else {
					direction = STEADY
				}
			}

			if (direction == DECREASING && diff < 0 && diff >= -3) || (direction == INCREASING && diff > 0 && diff <= 3) {
				isDiffInRange = true
			} else {
				isDiffInRange = false
			}

			if !isDiffInRange {
				ratings[i] = UNSAFE
				break
			}
		}
		if ratings[i] == UNSET {
			ratings[i] = SAFE
		}
	}

	for _, rating := range ratings {
		if rating == SAFE {
			count++
		}
	}

	return count
}
