package day8

type Vector struct {
	x, y int
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{v1.x - v2.x, v1.y - v2.y}
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func (v Vector) Negate() Vector {
	return Vector{-v.x, -v.y}
}

type Square struct {
	empty    bool
	antenna  rune
	antiNode bool
}

type Grid[T comparable] struct {
	values [][]T
	nrows  int
	ncols  int
}

func NewGrid[T comparable](nrows, ncols int) *Grid[T] {
	values := make([][]T, nrows)
	for i := range values {
		values[i] = make([]T, ncols)
	}
	return &Grid[T]{values, nrows, ncols}
}

func (g *Grid[T]) GetValue(v Vector) T {
	return g.values[g.nrows-v.y-1][v.x]
}

func (g *Grid[T]) SetValue(v Vector, value T) {
	g.values[g.nrows-v.y-1][v.x] = value
}

func (g *Grid[T]) InBounds(v Vector) bool {
	return v.x >= 0 && v.x < g.ncols && v.y >= 0 && v.y < g.nrows
}

func ParseGrid(inputs []string) (*Grid[Square], map[rune][]Vector) {
	nrows, ncols := len(inputs), len(inputs[0])
	g := NewGrid[Square](nrows, ncols)
	antennas := make(map[rune][]Vector)
	for i, input := range inputs {
		for j, r := range input {
			v := Vector{j, nrows - i - 1}
			if r == '.' {
				g.SetValue(v, Square{empty: true})
			} else {
				g.SetValue(v, Square{antenna: r})
				if antennas[r] == nil {
					antennas[r] = []Vector{v}
				} else {
					antennas[r] = append(antennas[r], v)
				}
			}
		}
	}
	return g, antennas
}

func CountAntiNodes(inputs []string) int {
	count := 0
	grid, antennas := ParseGrid(inputs)
	for _, vecs := range antennas {
		for i, v1 := range vecs {
			for j, v2 := range vecs[:i] {
				if i != j {
					diff := v1.Subtract(v2)
					antiNodes := []Vector{v1.Add(diff), v2.Add(diff.Negate())}
					for _, antiNode := range antiNodes {
						if grid.InBounds(antiNode) {
							value := grid.GetValue(antiNode)
							if !value.antiNode {
								grid.SetValue(antiNode, Square{
									empty:    value.empty,
									antenna:  value.antenna,
									antiNode: true,
								})
								count++
							}
						}
					}
				}
			}
		}
	}
	return count
}
