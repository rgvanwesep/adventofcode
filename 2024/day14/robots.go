package day14

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const robotPattern = `^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`

type vector struct {
	x, y int
}

func add(v1, v2 vector, nrows, ncols int) vector {
	x := v1.x + v2.x
	if x < 0 {
		x = x%ncols + ncols
	} else {
		x %= ncols
	}
	y := v1.y + v2.y
	if y < 0 {
		y = y%nrows + nrows
	} else {
		y %= nrows
	}
	return vector{x, y}
}

type robot struct {
	position, velocity vector
}

func (r *robot) update(nrows, ncols int) {
	r.position = add(r.position, r.velocity, nrows, ncols)
}

func parseRobots(inputs []string) []robot {
	robots := make([]robot, len(inputs))
	matcher := regexp.MustCompile(robotPattern)
	for i, s := range inputs {
		matched := matcher.FindStringSubmatch(s)

		x, err := strconv.Atoi(matched[1])
		if err != nil {
			log.Panicf("Could not parse position x from input: %q", s)
		}
		y, err := strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse position y from input: %q", s)
		}
		position := vector{x, y}

		x, err = strconv.Atoi(matched[3])
		if err != nil {
			log.Panicf("Could not parse velocity x from input: %q", s)
		}
		y, err = strconv.Atoi(matched[4])
		if err != nil {
			log.Panicf("Could not parse velocity y from input: %q", s)
		}
		velocity := vector{x, y}

		robots[i] = robot{position, velocity}

	}
	return robots
}

func CalcSafetyFactor(inputs []string, nrows, ncols int, nIter int) int {
	robots := parseRobots(inputs)
	for range nIter {
		for i := range robots {
			robots[i].update(nrows, ncols)
		}
	}
	sums := [4]int{}
	midCol := ncols / 2
	midRow := nrows / 2
	for _, r := range robots {
		switch {
		case r.position.x < midCol && r.position.y < midRow:
			sums[0]++
		case r.position.x > midCol && r.position.y < midRow:
			sums[1]++
		case r.position.x < midCol && r.position.y > midRow:
			sums[2]++
		case r.position.x > midCol && r.position.y > midRow:
			sums[3]++
		}
	}
	product := 1
	for _, n := range sums {
		product *= n
	}
	return product
}

func FindSignal(inputs []string, nrows, ncols int) int {
	robots := parseRobots(inputs)
	midCol := ncols / 2
	midRow := nrows / 2
	expectedSum := len(robots) * midCol * midRow / (ncols * nrows)
	threshold := 3 * expectedSum / 4
	i := 0
Outer:
	for {
		for i := range robots {
			robots[i].update(nrows, ncols)
		}
		sums := [4]int{}
		for _, r := range robots {
			switch {
			case r.position.x < midCol && r.position.y < midRow:
				sums[0]++
			case r.position.x > midCol && r.position.y < midRow:
				sums[1]++
			case r.position.x < midCol && r.position.y > midRow:
				sums[2]++
			case r.position.x > midCol && r.position.y > midRow:
				sums[3]++
			}
		}
		i++
		for _, n := range sums {
			diff := n - expectedSum
			if diff < -threshold || diff > threshold {
				break Outer
			}
		}
	}
	bitmap := make([][]bool, nrows)
	for i := range bitmap {
		bitmap[i] = make([]bool, ncols)
	}
	for _, r := range robots {
		bitmap[r.position.y][r.position.x] = true
	}
	for _, row := range bitmap {
		for _, value := range row {
			if value {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return i
}
