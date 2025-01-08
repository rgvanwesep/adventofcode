package day15

import (
	"aoc2024/deque"
	"iter"
)

const (
	emptyChar    = '.'
	robotChar    = '@'
	boxChar      = 'O'
	boxLeftChar  = '['
	boxRightChar = ']'
	wallChar     = '#'
	upChar       = '^'
	downChar     = 'v'
	leftChar     = '<'
	rightChar    = '>'
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

type set[T comparable] struct {
	values map[T]bool
}

func newSet[T comparable]() set[T] {
	return set[T]{
		values: map[T]bool{},
	}
}

func (s *set[T]) add(v T) {
	s.values[v] = true
}

func (s *set[T]) remove(v T) {
	delete(s.values, v)
}

func (s *set[T]) all() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.values {
			if !yield(v) {
				break
			}
		}
	}
}

func (s *set[T]) size() int {
	return len(s.values)
}

func (s *set[T]) clear() {
	s.values = map[T]bool{}
}

type wideWarehouse struct {
	grid        *grid
	robotPos    vector
	moves       []byte
	currentMove int
}

func newWideWarehouse(w warehouse) wideWarehouse {
	var robotPos vector
	wideGrid := newGrid(w.grid.nrows, 2*w.grid.ncols)
	for v, c := range w.grid.all() {
		wideVecs := [2]vector{{2 * v.x, v.y}, {2*v.x + 1, v.y}}
		switch c {
		case emptyChar:
			wideGrid.set(wideVecs[0], emptyChar)
			wideGrid.set(wideVecs[1], emptyChar)
		case robotChar:
			wideGrid.set(wideVecs[0], robotChar)
			wideGrid.set(wideVecs[1], emptyChar)
			robotPos = wideVecs[0]
		case boxChar:
			wideGrid.set(wideVecs[0], boxLeftChar)
			wideGrid.set(wideVecs[1], boxRightChar)
		case wallChar:
			wideGrid.set(wideVecs[0], wallChar)
			wideGrid.set(wideVecs[1], wallChar)
		}
	}
	return wideWarehouse{
		grid:        wideGrid,
		robotPos:    robotPos,
		moves:       w.moves,
		currentMove: 0,
	}
}

func (w *wideWarehouse) update() bool {
	if w.currentMove >= len(w.moves) {
		return false
	}
	updateVecs := newSet[vector]()
	frontVecs := deque.NewDeque[vector](-1)
	move := w.moves[w.currentMove]
	switch move {
	case upChar:
		updateVecs.add(vector{w.robotPos.x, w.robotPos.y})
		frontVecs.Append(vector{w.robotPos.x, w.robotPos.y + 1})
	frontLoopUp:
		for {
			if frontVec, ok := frontVecs.Pop(); ok {
				c := w.grid.get(frontVec)
				switch c {
				case boxLeftChar:
					updateVecs.add(vector{frontVec.x, frontVec.y})
					updateVecs.add(vector{frontVec.x + 1, frontVec.y})
					frontVecs.Append(vector{frontVec.x, frontVec.y + 1})
					frontVecs.Append(vector{frontVec.x + 1, frontVec.y + 1})
				case boxRightChar:
					updateVecs.add(vector{frontVec.x, frontVec.y})
					updateVecs.add(vector{frontVec.x - 1, frontVec.y})
					frontVecs.Append(vector{frontVec.x, frontVec.y + 1})
					frontVecs.Append(vector{frontVec.x - 1, frontVec.y + 1})
				case wallChar:
					updateVecs.clear()
					break frontLoopUp
				}
			} else {
				break frontLoopUp
			}
		}
		for updateVecs.size() > 0 {
			for updateVec := range updateVecs.all() {
				nextVec := vector{updateVec.x, updateVec.y + 1}
				if w.grid.get(nextVec) == emptyChar {
					c := w.grid.get(updateVec)
					w.grid.set(nextVec, c)
					w.grid.set(updateVec, emptyChar)
					if c == robotChar {
						w.robotPos = nextVec
					}
					updateVecs.remove(updateVec)
				}
			}
		}
	case downChar:
		updateVecs.add(vector{w.robotPos.x, w.robotPos.y})
		frontVecs.Append(vector{w.robotPos.x, w.robotPos.y - 1})
	frontLoopDown:
		for {
			if frontVec, ok := frontVecs.Pop(); ok {
				c := w.grid.get(frontVec)
				switch c {
				case boxLeftChar:
					updateVecs.add(vector{frontVec.x, frontVec.y})
					updateVecs.add(vector{frontVec.x + 1, frontVec.y})
					frontVecs.Append(vector{frontVec.x, frontVec.y - 1})
					frontVecs.Append(vector{frontVec.x + 1, frontVec.y - 1})
				case boxRightChar:
					updateVecs.add(vector{frontVec.x, frontVec.y})
					updateVecs.add(vector{frontVec.x - 1, frontVec.y})
					frontVecs.Append(vector{frontVec.x, frontVec.y - 1})
					frontVecs.Append(vector{frontVec.x - 1, frontVec.y - 1})
				case wallChar:
					updateVecs.clear()
					break frontLoopDown
				}
			} else {
				break frontLoopDown
			}
		}
		for updateVecs.size() > 0 {
			for updateVec := range updateVecs.all() {
				nextVec := vector{updateVec.x, updateVec.y - 1}
				if w.grid.get(nextVec) == emptyChar {
					c := w.grid.get(updateVec)
					w.grid.set(nextVec, c)
					w.grid.set(updateVec, emptyChar)
					if c == robotChar {
						w.robotPos = nextVec
					}
					updateVecs.remove(updateVec)
				}
			}
		}
	case leftChar:
		updateVecs.add(vector{w.robotPos.x, w.robotPos.y})
		frontVecs.Append(vector{w.robotPos.x - 1, w.robotPos.y})
	frontLoopLeft:
		for {
			if frontVec, ok := frontVecs.Pop(); ok {
				c := w.grid.get(frontVec)
				switch c {
				case boxRightChar:
					updateVecs.add(vector{frontVec.x, frontVec.y})
					updateVecs.add(vector{frontVec.x - 1, frontVec.y})
					frontVecs.Append(vector{frontVec.x - 2, frontVec.y})
				case wallChar:
					updateVecs.clear()
					break frontLoopLeft
				}
			} else {
				break frontLoopLeft
			}
		}
		for updateVecs.size() > 0 {
			for updateVec := range updateVecs.all() {
				nextVec := vector{updateVec.x - 1, updateVec.y}
				if w.grid.get(nextVec) == emptyChar {
					c := w.grid.get(updateVec)
					w.grid.set(nextVec, c)
					w.grid.set(updateVec, emptyChar)
					if c == robotChar {
						w.robotPos = nextVec
					}
					updateVecs.remove(updateVec)
				}
			}
		}
	case rightChar:
		updateVecs.add(vector{w.robotPos.x, w.robotPos.y})
		frontVecs.Append(vector{w.robotPos.x + 1, w.robotPos.y})
	frontLoopRight:
		for {
			if frontVec, ok := frontVecs.Pop(); ok {
				c := w.grid.get(frontVec)
				switch c {
				case boxLeftChar:
					updateVecs.add(vector{frontVec.x, frontVec.y})
					updateVecs.add(vector{frontVec.x + 1, frontVec.y})
					frontVecs.Append(vector{frontVec.x + 2, frontVec.y})
				case wallChar:
					updateVecs.clear()
					break frontLoopRight
				}
			} else {
				break frontLoopRight
			}
		}
		for updateVecs.size() > 0 {
			for updateVec := range updateVecs.all() {
				nextVec := vector{updateVec.x + 1, updateVec.y}
				if w.grid.get(nextVec) == emptyChar {
					c := w.grid.get(updateVec)
					w.grid.set(nextVec, c)
					w.grid.set(updateVec, emptyChar)
					if c == robotChar {
						w.robotPos = nextVec
					}
					updateVecs.remove(updateVec)
				}
			}
		}
	}
	w.currentMove++
	return true
}

func (w *wideWarehouse) sumCoordinates() int {
	sum := 0
	for v, c := range w.grid.all() {
		if c == boxLeftChar {
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

func SumCoordinatesWide(inputs []string) int {
	w := newWideWarehouse(parseWarehouse(inputs))
	for w.update() {
	}
	return w.sumCoordinates()
}
