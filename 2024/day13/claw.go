package day13

import (
	"log"
	"regexp"
	"strconv"
)

const (
	aCost         = 3
	bCost         = 1
	buttonPattern = `^Button (A|B): X\+(\d+), Y\+(\d+)$`
	prizePattern  = `^Prize: X=(\d+), Y=(\d+)$`
	maxUInt       = ^uint(0)
	maxInt        = int(maxUInt >> 1)
)

type vector struct {
	x, y int
}

func add(v1, v2 vector) vector {
	return vector{v1.x + v2.x, v1.y + v2.y}
}

type grid struct {
	xMax, yMax int
	values     []int
}

func newGrid(xMax, yMax int) *grid {
	size := xMax * yMax
	values := make([]int, size)
	for i := range size {
		values[i] = maxInt
	}
	return &grid{xMax, yMax, values}
}

func (g *grid) get(v vector) (int, bool) {
	i := v.y*g.xMax + v.x
	if i < len(g.values) {
		c := g.values[i]
		if c != maxInt {
			return g.values[i], true
		}
	}
	return 0, false
}

func (g *grid) set(v vector, c int) bool {
	i := v.y*g.xMax + v.x
	if i < len(g.values) {
		g.values[i] = c
		return true
	}
	return false
}

type machine struct {
	aVec, bVec, prizeVec vector
	memo                 *grid
}

func (m *machine) minCost(origin vector) int {
	if c, ok := m.memo.get(origin); ok {
		return c
	}
	if origin.x == m.prizeVec.x && origin.y == m.prizeVec.y {
		m.memo.set(origin, 0)
		return 0
	}
	if origin.x > m.prizeVec.x || origin.y > m.prizeVec.y {
		return maxInt
	}
	aNextOrigin := add(origin, m.aVec)
	aNextCost := m.minCost(aNextOrigin)
	bNextOrigin := add(origin, m.bVec)
	bNextCost := m.minCost(bNextOrigin)
	if aNextCost == maxInt && bNextCost == maxInt {
		return maxInt
	}
	if aNextCost == maxInt {
		return bCost + bNextCost
	}
	if bNextCost == maxInt {
		return aCost + aNextCost
	}
	result := min(aCost+aNextCost, bCost+bNextCost)
	m.memo.set(origin, result)
	return result
}

func parseMachines(inputs []string) []*machine {
	buttonMatcher := regexp.MustCompile(buttonPattern)
	prizeMatcher := regexp.MustCompile(prizePattern)
	machines := []*machine{}
	for i := 0; i < len(inputs); i += 4 {
		matched := buttonMatcher.FindStringSubmatch(inputs[i])
		x, err := strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse Button A x value from input %q", inputs[i])
		}
		y, err := strconv.Atoi(matched[3])
		if err != nil {
			log.Panicf("Could not parse Button A y value from input %q", inputs[i])
		}
		aVec := vector{x, y}

		matched = buttonMatcher.FindStringSubmatch(inputs[i+1])
		x, err = strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse Button B x value from input %q", inputs[i+1])
		}
		y, err = strconv.Atoi(matched[3])
		if err != nil {
			log.Panicf("Could not parse Button B y value from input %q", inputs[i+1])
		}
		bVec := vector{x, y}

		matched = prizeMatcher.FindStringSubmatch(inputs[i+2])
		x, err = strconv.Atoi(matched[1])
		if err != nil {
			log.Panicf("Could not parse Prize x value from input %q", inputs[i+2])
		}
		y, err = strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse Prize B y value from input %q", inputs[i+2])
		}
		prizeVec := vector{x, y}

		machines = append(machines, &machine{aVec, bVec, prizeVec, newGrid(prizeVec.x, prizeVec.y)})
	}
	return machines
}

func MinCost(inputs []string) int {
	sum := 0
	origin := vector{0, 0}
	machines := parseMachines(inputs)
	for _, m := range machines {
		if c := m.minCost(origin); c != maxInt {
			sum += c
		}
	}
	return sum
}
