package day15

import "iter"

const (
	emptyChar = '.'
	robotChar = '@'
	boxChar   = 'O'
	wallChar  = '#'
	upChar    = '^'
	downChar  = 'v'
	leftChar  = '<'
	rightChar = '>'
)

type vector struct {
	x, y int
}

type grid struct {
	nrows, ncols int
	values       []byte
}

func newGrid(nrows, ncols int) *grid {
	return &grid{
		nrows:  nrows,
		ncols:  ncols,
		values: make([]byte, nrows*ncols),
	}
}

func (g *grid) get(v vector) byte {
	row := g.nrows - v.y - 1
	return g.values[row*g.ncols+v.x]
}

func (g *grid) set(v vector, c byte) {
	row := g.nrows - v.y - 1
	g.values[row*g.ncols+v.x] = c
}

func (g *grid) all() iter.Seq2[vector, byte] {
	return func(yield func(vector, byte) bool) {
		for i, c := range g.values {
			row := i / g.ncols
			col := i % g.ncols
			v := vector{col, g.nrows - row - 1}
			if !yield(v, c) {
				break
			}
		}
	}
}

type warehouse struct {
	grid        *grid
	robotPos    vector
	moves       []byte
	currentMove int
}

func parseWarehouse(inputs []string) warehouse {
	gridRows := []string{}
	moveRows := []string{}
	inGrid := true
	for _, row := range inputs {
		if row == "" {
			inGrid = false
			continue
		}
		if inGrid {
			gridRows = append(gridRows, row)
		} else {
			moveRows = append(moveRows, row)
		}
	}
	nrows := len(gridRows)
	ncols := len(gridRows[0])
	g := newGrid(nrows, ncols)
	var robotPos vector
	for i, row := range gridRows {
		for j := range row {
			v := vector{j, nrows - i - 1}
			c := row[j]
			g.set(v, c)
			if c == robotChar {
				robotPos = v
			}
		}
	}
	moves := []byte{}
	for _, row := range moveRows {
		for i := range row {
			moves = append(moves, row[i])
		}
	}
	return warehouse{
		grid:        g,
		robotPos:    robotPos,
		moves:       moves,
		currentMove: 0,
	}
}

func (w *warehouse) update() bool {
	if w.currentMove >= len(w.moves) {
		return false
	}
	move := w.moves[w.currentMove]
	switch move {
	case upChar:
		for y := w.robotPos.y + 1; y < w.grid.nrows; y++ {
			c := w.grid.get(vector{w.robotPos.x, y})
			if c == emptyChar {
				newRobotPos := vector{w.robotPos.x, w.robotPos.y + 1}
				w.grid.set(vector{w.robotPos.x, y}, boxChar)
				w.grid.set(newRobotPos, robotChar)
				w.grid.set(vector{w.robotPos.x, w.robotPos.y}, emptyChar)
				w.robotPos = newRobotPos
				break
			} else if c == wallChar {
				break
			}
		}
	case downChar:
		for y := w.robotPos.y - 1; y >= 0; y-- {
			c := w.grid.get(vector{w.robotPos.x, y})
			if c == emptyChar {
				newRobotPos := vector{w.robotPos.x, w.robotPos.y - 1}
				w.grid.set(vector{w.robotPos.x, y}, boxChar)
				w.grid.set(newRobotPos, robotChar)
				w.grid.set(vector{w.robotPos.x, w.robotPos.y}, emptyChar)
				w.robotPos = newRobotPos
				break
			} else if c == wallChar {
				break
			}
		}
	case leftChar:
		for x := w.robotPos.x - 1; x >= 0; x-- {
			c := w.grid.get(vector{x, w.robotPos.y})
			if c == emptyChar {
				newRobotPos := vector{w.robotPos.x - 1, w.robotPos.y}
				w.grid.set(vector{x, w.robotPos.y}, boxChar)
				w.grid.set(newRobotPos, robotChar)
				w.grid.set(vector{w.robotPos.x, w.robotPos.y}, emptyChar)
				w.robotPos = newRobotPos
				break
			} else if c == wallChar {
				break
			}
		}
	case rightChar:
		for x := w.robotPos.x + 1; x < w.grid.ncols; x++ {
			c := w.grid.get(vector{x, w.robotPos.y})
			if c == emptyChar {
				newRobotPos := vector{w.robotPos.x + 1, w.robotPos.y}
				w.grid.set(vector{x, w.robotPos.y}, boxChar)
				w.grid.set(newRobotPos, robotChar)
				w.grid.set(vector{w.robotPos.x, w.robotPos.y}, emptyChar)
				w.robotPos = newRobotPos
				break
			} else if c == wallChar {
				break
			}
		}
	}
	w.currentMove++
	return true
}

func (w *warehouse) sumCoordinates() int {
	sum := 0
	for v, c := range w.grid.all() {
		if c == boxChar {
			sum += 100*(w.grid.nrows-v.y-1) + v.x
		}
	}
	return sum
}

func SumCoordinates(inputs []string) int {
	w := parseWarehouse(inputs)
	for w.update() {
	}
	return w.sumCoordinates()
}
