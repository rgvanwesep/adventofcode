package day25

import "iter"

const (
	emptyChar = '.'
	fillChar  = '#'
)

type vector struct {
	x, y int
}

type grid struct {
	nrows, ncols int
	values       []byte
}

func newGrid(nrows, ncols int) *grid {
	values := make([]byte, nrows*ncols)
	return &grid{
		nrows:  nrows,
		ncols:  ncols,
		values: values,
	}
}

func (g *grid) get(v vector) (byte, bool) {
	if v.x < 0 || v.x >= g.ncols || v.y < 0 || v.y >= g.nrows {
		return 0, false
	}
	return g.values[v.y*g.ncols+v.x], true
}

func (g *grid) set(v vector, c byte) {
	g.values[v.y*g.ncols+v.x] = c
}

func (g *grid) all() iter.Seq2[vector, byte] {
	return func(yield func(vector, byte) bool) {
		for i, c := range g.values {
			row := i / g.ncols
			col := i % g.ncols
			v := vector{col, row}
			if !yield(v, c) {
				break
			}
		}
	}
}

func (g grid) String() string {
	s := []byte{}
	for i, c := range g.values {
		s = append(s, c)
		if i%g.ncols == g.ncols-1 {
			s = append(s, '\n')
		}
	}
	return string(s)
}

type key struct {
	grid    *grid
	heights []int
}

func newKey(g *grid) key {
	heights := make([]int, g.ncols)
	for x := range g.ncols {
		for y := g.nrows - 2; y >= 0; y-- {
			if c, _ := g.get(vector{x, y}); c == fillChar {
				heights[x]++
			} else {
				break
			}
		}
	}
	return key{g, heights}
}

type lock struct {
	grid    *grid
	heights []int
}

func newLock(g *grid) lock {
	heights := make([]int, g.ncols)
	for x := range g.ncols {
		for y := 1; y < g.nrows; y++ {
			if c, _ := g.get(vector{x, y}); c == fillChar {
				heights[x]++
			} else {
				break
			}
		}
	}
	return lock{g, heights}
}

func doesFit(k key, l lock) bool {
	for i, h := range k.heights {
		if h+l.heights[i] >= k.grid.nrows-1 {
			return false
		}
	}
	return true
}

func parseKeysAndLocks(inputs []string) ([]key, []lock) {
	var inKey, inNewGrid bool
	ncols := len(inputs[0])
	nrows := ncols + 2
	keys := []key{}
	locks := []lock{}
	g := newGrid(nrows, ncols)
	inNewGrid = true
	i := 0
	for _, input := range inputs {
		if input == "" {
			if inKey {
				keys = append(keys, newKey(g))
			} else {
				locks = append(locks, newLock(g))
			}
			g = newGrid(nrows, ncols)
			inNewGrid = true
			i = 0
			continue
		}
		if inNewGrid {
			inKey = input[0] == emptyChar
			inNewGrid = false
		}
		for j := range input {
			g.set(vector{j, i}, input[j])
		}
		i++
	}
	if inKey {
		keys = append(keys, newKey(g))
	} else {
		locks = append(locks, newLock(g))
	}
	return keys, locks
}

func CountFits(inputs []string) int {
	count := 0
	keys, locks := parseKeysAndLocks(inputs)
	for _, k := range keys {
		for _, l := range locks {
			if doesFit(k, l) {
				count++
			}
		}
	}
	return count
}
