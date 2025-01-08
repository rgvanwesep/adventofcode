package day10

import (
	"aoc2024/deque"
	"log"
)

const ZERO byte = '0'

type Point struct {
	x, y int
}

type Grid[T any] struct {
	nrows, ncols int
	values       []T
}

func NewGrid[T any](nrows, ncols int) *Grid[T] {
	return &Grid[T]{
		nrows:  nrows,
		ncols:  ncols,
		values: make([]T, nrows*ncols),
	}
}

func (g *Grid[T]) GetValue(p Point) (v T, ok bool) {
	if p.x >= 0 && p.x < g.ncols && p.y >= 0 && p.y < g.nrows {
		row := g.nrows - p.y - 1
		v = g.values[g.ncols*row+p.x]
		ok = true
	}
	return
}

func (g *Grid[T]) SetValue(p Point, v T) (ok bool) {
	if p.x >= 0 && p.x < g.ncols && p.y >= 0 && p.y < g.nrows {
		row := g.nrows - p.y - 1
		g.values[g.ncols*row+p.x] = v
		ok = true
	}
	return
}

type Hiker struct {
	origin, position   Point
	halted, reachedTop bool
}

func (h Hiker) Step(g *Grid[byte]) []Hiker {
	height, _ := g.GetValue(h.position)
	if height == 9 {
		return []Hiker{{
			origin:     h.origin,
			position:   h.position,
			halted:     true,
			reachedTop: true,
		}}
	}
	neighbors := []Point{
		{h.position.x - 1, h.position.y},
		{h.position.x + 1, h.position.y},
		{h.position.x, h.position.y - 1},
		{h.position.x, h.position.y + 1},
	}
	hikers := []Hiker{}
	for _, neighbor := range neighbors {
		if neighborHeight, ok := g.GetValue(neighbor); ok {
			if neighborHeight-height == 1 {
				hikers = append(hikers, Hiker{
					origin:   h.origin,
					position: neighbor,
				})
			}
		}
	}
	if len(hikers) == 0 {
		return []Hiker{{
			origin:     h.origin,
			position:   h.position,
			halted:     true,
			reachedTop: false,
		}}
	}
	return hikers
}

func ParseGrid(inputs []string) (*Grid[byte], []Point) {
	nrows, ncols := len(inputs), len(inputs[0])
	grid := NewGrid[byte](nrows, ncols)
	trailHeads := []Point{}
	for row, input := range inputs {
		for col := range input {
			p := Point{col, nrows - row - 1}
			value := input[col] - ZERO
			if ok := grid.SetValue(p, value); !ok {
				log.Panicf("Could not set value for %v", p)
			}
			if value == 0 {
				trailHeads = append(trailHeads, p)
			}
		}
	}
	return grid, trailHeads
}

func RunHikers(grid *Grid[byte], trailHeads []Point) deque.Deque[Hiker] {
	hikers := deque.NewDeque[Hiker](-1)
	for _, p := range trailHeads {
		hikers.Append(Hiker{
			origin:   p,
			position: p,
		})
	}
	halted := deque.NewDeque[Hiker](-1)
	for {
		if hiker, ok := hikers.Pop(); ok {
			newHikers := hiker.Step(grid)
			for _, h := range newHikers {
				if h.halted {
					halted.Append(h)
				} else {
					hikers.Append(h)
				}
			}
		} else {
			break
		}
	}
	return halted
}

func SumTrailScores(inputs []string) int {
	grid, trailHeads := ParseGrid(inputs)
	halted := RunHikers(grid, trailHeads)
	pairs := map[[2]Point]bool{}
	for {
		if hiker, ok := halted.Pop(); ok {
			if hiker.reachedTop {
				pairs[[2]Point{hiker.origin, hiker.position}] = true
			}
		} else {
			break
		}
	}
	return len(pairs)
}

func SumTrailRatings(inputs []string) int {
	grid, trailHeads := ParseGrid(inputs)
	halted := RunHikers(grid, trailHeads)
	count := 0
	for {
		if hiker, ok := halted.Pop(); ok {
			if hiker.reachedTop {
				count++
			}
		} else {
			break
		}
	}
	return count
}
