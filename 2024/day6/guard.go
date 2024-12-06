package day6

import (
	"fmt"
	"maps"
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
	direction rune
}

func (g Guard) Turned() Guard {
	var gTurned Guard
	switch g.direction {
	case UP:
		gTurned = Guard{RIGHT}
	case DOWN:
		gTurned = Guard{LEFT}
	case LEFT:
		gTurned = Guard{UP}
	case RIGHT:
		gTurned = Guard{DOWN}
	}
	return gTurned
}

type Coordinates struct {
	x, y int
}

type Square struct {
	guard    Optional[Guard]
	occupant rune
	visted   bool
}

type Grid struct {
	nrows, ncols      int
	nguards, nvisited int
	squares           map[Coordinates]Square
	guardCoords       []Coordinates
}

func NewGrid(nrows, ncols int) Grid {
	squares := make(map[Coordinates]Square)
	return Grid{nrows, ncols, 0, 0, squares, []Coordinates{}}
}

func (g *Grid) AddSquare(c Coordinates, s Square) {
	g.squares[c] = s
	if !s.guard.IsNone() {
		g.nguards++
		g.nvisited++
		g.guardCoords = append(g.guardCoords, c)
	}
}

func (g *Grid) Step() {
	var (
		cNext Coordinates
		sNext Square
		gNext Guard
	)
	updates := make(map[Coordinates]Square)
	newGuardCoords := make([]Coordinates, 0)
	for _, c := range g.guardCoords {
		s := g.squares[c]
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
			sNext = g.squares[cNext]
			if sNext.occupant == OBSTRUCTION {
				gNext = s.guard.GetValue().Turned()
				updates[c] = Square{Some(gNext), s.occupant, s.visted}
				newGuardCoords = append(newGuardCoords, c)
			} else {
				gNext = s.guard.GetValue()
				updates[c] = Square{None[Guard](), s.occupant, s.visted}
				updates[cNext] = Square{Some(gNext), sNext.occupant, true}
				newGuardCoords = append(newGuardCoords, cNext)
				if !g.squares[cNext].visted {
					g.nvisited++
				}
			}
		}
	}
	maps.Insert(g.squares, maps.All(updates))
	g.guardCoords = newGuardCoords
}

func (g Grid) String() string {
	runes := make([]rune, 0)
	for i := range g.nrows {
		for j := range g.ncols {
			c := Coordinates{j, g.nrows - i - 1}
			s := g.squares[c]
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
			case UP, DOWN, LEFT, RIGHT:
				s = Square{Some(Guard{char}), EMPTY, true}
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
