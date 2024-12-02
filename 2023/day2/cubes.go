package day2

import (
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type Set struct {
	Red   int
	Green int
	Blue  int
}

func NewSet(input string) Set {
	counts := strings.Split(input, ", ")
	countMap := make(map[string]int)
	for _, count := range counts {
		pair := strings.Split(count, " ")
		countMap[pair[1]], _ = strconv.Atoi(pair[0])
	}
	return Set{countMap["red"], countMap["green"], countMap["blue"]}
}

func (set Set) Check() bool {
	return set.Red <= maxRed && set.Green <= maxGreen && set.Blue <= maxBlue
}

func (set Set) Power() int {
	return set.Red * set.Green * set.Blue
}

func (set Set) MergeMax(other Set) Set {
	return Set{
		max(set.Red, other.Red),
		max(set.Green, other.Green),
		max(set.Blue, other.Blue),
	}
}

type Game struct {
	Id   int
	Sets []Set
}

func NewGame(input string) Game {
	input = strings.TrimRight(input, "\n")
	parts := strings.Split(input, ": ")

	gamePair := strings.Split(parts[0], " ")
	id, _ := strconv.Atoi(gamePair[1])

	sets := make([]Set, 0)
	for _, set := range strings.Split(parts[1], "; ") {
		sets = append(sets, NewSet(set))
	}
	return Game{id, sets}
}

func (game Game) Check() bool {
	for _, set := range game.Sets {
		if !set.Check() {
			return false
		}
	}
	return true
}

func (game Game) LeastUpperBound() Set {
	bound := Set{}
	for _, set := range game.Sets {
		bound = bound.MergeMax(set)
	}
	return bound
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		if game := NewGame(input); game.Check() {
			sum += game.Id
		}
	}
	return sum
}

func SumPower(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		game := NewGame(input)
		sum += game.LeastUpperBound().Power()
	}
	return sum
}
