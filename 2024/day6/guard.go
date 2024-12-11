package day6

import (
	"fmt"
	"log"
)

const (
	UP          = '^'
	DOWN        = 'v'
	LEFT        = '<'
	RIGHT       = '>'
	EMPTY       = '.'
	OBSTRUCTION = '#'
)

type Optional[T any] []T

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{value}
}

func (o Optional[T]) IsNone() bool {
	return len(o) == 0
}

func (o Optional[T]) GetValue() T {
	if o.IsNone() {
		panic("Cannot get value of None")
	}
	return o[0]
}

type Guard struct {
	direction      rune
	directionIndex int
}

func (g Guard) Turned() Guard {
	var gTurned Guard
	switch g.direction {
	case UP:
		gTurned = Guard{RIGHT, 3}
	case DOWN:
		gTurned = Guard{LEFT, 2}
	case LEFT:
		gTurned = Guard{UP, 0}
	case RIGHT:
		gTurned = Guard{DOWN, 1}
	}
	return gTurned
}

type Coordinates struct {
	x, y int
}

func (c Coordinates) RowMajorIndex(nrows, ncols int) int {
	row := nrows - c.y - 1
	return row*ncols + c.x
}

type Square struct {
	guard    Optional[Guard]
	occupant rune
	visted   bool
}

type Grid struct {
	nrows, ncols      int
	nguards, nvisited int
	squares           []Square
	guardCoords       Coordinates
	stateCounts       []int
}

func NewGrid(nrows, ncols int) Grid {
	squares := make([]Square, nrows*ncols)
	guardCoords := make([]int, nrows*ncols*4)
	return Grid{nrows, ncols, 0, 0, squares, Coordinates{-1, -1}, guardCoords}
}

func (g *Grid) GetValue(c Coordinates) Square {
	return g.squares[c.RowMajorIndex(g.nrows, g.ncols)]
}

func (g *Grid) SetValue(c Coordinates, s Square) {
	g.squares[c.RowMajorIndex(g.nrows, g.ncols)] = s
}

func (g *Grid) GetStateCount() int {
	guard := g.GetValue(g.guardCoords).guard
	if guard.IsNone() {
		return -1
	}
	row := g.nrows - g.guardCoords.y - 1
	col := g.guardCoords.x
	directionIndex := guard.GetValue().directionIndex
	return g.stateCounts[row*g.ncols*4+col*4+directionIndex]
}

func (g *Grid) IncrementStateCount() {
	guard := g.GetValue(g.guardCoords).guard
	if guard.IsNone() {
		return
	}
	row := g.nrows - g.guardCoords.y - 1
	col := g.guardCoords.x
	directionIndex := guard.GetValue().directionIndex
	g.stateCounts[row*g.ncols*4+col*4+directionIndex]++
}

func (g *Grid) AddSquare(c Coordinates, s Square) {
	g.SetValue(c, s)
	if !s.guard.IsNone() {
		g.nguards++
		g.nvisited++
		g.guardCoords = c
	}
}

func (g *Grid) Step() {
	var (
		cNext Coordinates
		sNext Square
		gNext Guard
	)
	updates := make(map[Coordinates]Square)
	newGuardCoords := Coordinates{-1, -1}
	c := g.guardCoords

	s := g.GetValue(c)
	switch s.guard.GetValue().direction {
	case UP:
		cNext = Coordinates{c.x, c.y + 1}
	case DOWN:
		cNext = Coordinates{c.x, c.y - 1}
	case LEFT:
		cNext = Coordinates{c.x - 1, c.y}
	case RIGHT:
		cNext = Coordinates{c.x + 1, c.y}
	}
	if cNext.x < 0 || cNext.x >= g.ncols || cNext.y < 0 || cNext.y >= g.nrows {
		updates[c] = Square{None[Guard](), s.occupant, s.visted}
		g.nguards--
	} else {
		sNext = g.GetValue(cNext)
		if sNext.occupant == OBSTRUCTION {
			gNext = s.guard.GetValue().Turned()
			updates[c] = Square{Some(gNext), s.occupant, s.visted}
			newGuardCoords = c
		} else {
			gNext = s.guard.GetValue()
			updates[c] = Square{None[Guard](), s.occupant, s.visted}
			updates[cNext] = Square{Some(gNext), sNext.occupant, true}
			newGuardCoords = cNext
			if !g.GetValue(cNext).visted {
				g.nvisited++
			}
		}
	}

	for c, s := range updates {
		g.SetValue(c, s)
	}
	g.guardCoords = newGuardCoords
}

func (g Grid) String() string {
	runes := make([]rune, 0)
	for i := range g.nrows {
		for j := range g.ncols {
			c := Coordinates{j, g.nrows - i - 1}
			s := g.GetValue(c)
			if s.guard.IsNone() {
				runes = append(runes, s.occupant)
			} else {
				runes = append(runes, s.guard.GetValue().direction)
			}
		}
		runes = append(runes, '\n')
	}
	return string(runes)
}

func ParseGrid(inputs []string) Grid {
	nrows, ncols := len(inputs), len(inputs[0])
	grid := NewGrid(nrows, ncols)
	for i, row := range inputs {
		for j, char := range row {
			c := Coordinates{j, nrows - i - 1}
			var s Square
			switch char {
			case UP:
				s = Square{Some(Guard{char, 0}), EMPTY, true}
			case DOWN:
				s = Square{Some(Guard{char, 1}), EMPTY, true}
			case LEFT:
				s = Square{Some(Guard{char, 2}), EMPTY, true}
			case RIGHT:
				s = Square{Some(Guard{char, 3}), EMPTY, true}
			case EMPTY:
				s = Square{None[Guard](), EMPTY, false}
			case OBSTRUCTION:
				s = Square{None[Guard](), OBSTRUCTION, false}
			default:
				panic(fmt.Sprintf("Invalid character at (%d, %d): %q", i, j, char))
			}
			grid.AddSquare(c, s)
		}
	}
	return grid
}

func CountVisited(inputs []string) int {
	grid := ParseGrid(inputs)
	for grid.nguards > 0 {
		grid.Step()
	}
	return grid.nvisited
}

func InputsToRunes(inputs []string) [][]rune {
	runes := make([][]rune, len(inputs))
	for i, input := range inputs {
		runes[i] = make([]rune, len(input))
		for j, r := range input {
			runes[i][j] = r
		}
	}
	return runes
}

func RunesToInputs(runes [][]rune) []string {
	inputs := make([]string, len(runes))
	for i, row := range runes {
		inputs[i] = string(row)
	}
	return inputs
}

func CountCyclingObstructions(inputs []string) int {
	count := 0
	variations := make([][]string, 0)
	for i := range len(inputs) {
		for j := range len(inputs[0]) {
			runes := InputsToRunes(inputs)
			if runes[i][j] == EMPTY {
				runes[i][j] = OBSTRUCTION
				variations = append(variations, RunesToInputs(runes))
			}
		}
	}
	log.Printf("Running with %d variations", len(variations))
	for i, variation := range variations {
		grid := ParseGrid(variation)
		grid.IncrementStateCount()
		for grid.nguards > 0 {
			grid.Step()
			grid.IncrementStateCount()
			if grid.GetStateCount() > 1 {
				count++
				break
			}
		}
		log.Printf("Count is %d at variation %d", count, i)
	}
	return count
}
