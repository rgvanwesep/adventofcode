package day4

import "fmt"

const (
	X, M, A, S byte = 'X', 'M', 'A', 'S'
	UP              = iota
	DOWN
	LEFT
	RIGHT
	UP_LEFT
	UP_RIGHT
	DOWN_LEFT
	DOWN_RIGHT
)

func ParseGrid(inputs []string) [][]byte {
	var (
		grid [][]byte
	)
	grid = make([][]byte, 0)
	for _, rowStr := range inputs {
		row := make([]byte, 0)
		for _, char := range rowStr {
			row = append(row, byte(char))
		}
		grid = append(grid, row)
	}
	return grid
}

func Neighbors(pos [2]int, nrows int, ncols int) [][2]int {
	var (
		neighbors [][2]int
		x, y      int
	)
	x, y = pos[0], pos[1]
	maybeNeighbors := [][2]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
		{x - 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y - 1},
		{x + 1, y + 1},
	}
	neighbors = make([][2]int, 0)
	for _, maybeNeighbor := range maybeNeighbors {
		x, y = maybeNeighbor[0], maybeNeighbor[1]
		if x >= 0 && x < nrows && y >= 0 && y < ncols {
			neighbors = append(neighbors, maybeNeighbor)
		}
	}
	return neighbors
}

func Diagonals(pos [2]int, nrows int, ncols int) [][2]int {
	var (
		diagonals [][2]int
		x, y      int
	)
	x, y = pos[0], pos[1]
	/*
	 * 0.2
	 * ...
	 * 1.3
	 */
	maybeDiagonals := [][2]int{
		{x - 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y - 1},
		{x + 1, y + 1},
	}
	diagonals = make([][2]int, 0)
	for _, maybeNeighbor := range maybeDiagonals {
		x, y = maybeNeighbor[0], maybeNeighbor[1]
		if x >= 0 && x < nrows && y >= 0 && y < ncols {
			diagonals = append(diagonals, maybeNeighbor)
		}
	}
	return diagonals
}

func FindEnd(grid [][]byte, posM [2]int, direction int) [][2]int {
	var (
		end           [2]int
		nextPositions [2][2]int
	)
	x, y := posM[0], posM[1]
	nrows, ncols := len(grid), len(grid[0])
	switch direction {
	case UP:
		if y >= 2 {
			nextPositions = [2][2]int{{x, y - 1}, {x, y - 2}}
		} else {
			return [][2]int{}
		}
	case DOWN:
		if y < nrows-2 {
			nextPositions = [2][2]int{{x, y + 1}, {x, y + 2}}
		} else {
			return [][2]int{}
		}
	case LEFT:
		if x >= 2 {
			nextPositions = [2][2]int{{x - 1, y}, {x - 2, y}}
		} else {
			return [][2]int{}
		}
	case RIGHT:
		if x < ncols-2 {
			nextPositions = [2][2]int{{x + 1, y}, {x + 2, y}}
		} else {
			return [][2]int{}
		}
	case UP_LEFT:
		if x >= 2 && y >= 2 {
			nextPositions = [2][2]int{{x - 1, y - 1}, {x - 2, y - 2}}
		} else {
			return [][2]int{}
		}
	case UP_RIGHT:
		if y >= 2 && x < ncols-2 {
			nextPositions = [2][2]int{{x + 1, y - 1}, {x + 2, y - 2}}
		} else {
			return [][2]int{}
		}
	case DOWN_LEFT:
		if y < nrows-2 && x >= 2 {
			nextPositions = [2][2]int{{x - 1, y + 1}, {x - 2, y + 2}}
		} else {
			return [][2]int{}
		}
	case DOWN_RIGHT:
		if x < nrows-2 && y < ncols-2 {
			nextPositions = [2][2]int{{x + 1, y + 1}, {x + 2, y + 2}}
		} else {
			return [][2]int{}
		}
	default:
		panic(fmt.Sprintf("Invalid direction: %v", direction))
	}
	maybeA := grid[nextPositions[0][1]][nextPositions[0][0]]
	maybeS := grid[nextPositions[1][1]][nextPositions[1][0]]
	if maybeA == A && maybeS == S {
		end = nextPositions[1]
	} else {
		return [][2]int{}
	}
	return [][2]int{end}
}

func FindDirection(start, end [2]int) int {
	var direction int
	switch diff := [2]int{end[0] - start[0], end[1] - start[1]}; diff {
	case [2]int{0, -1}:
		direction = UP
	case [2]int{0, 1}:
		direction = DOWN
	case [2]int{-1, 0}:
		direction = LEFT
	case [2]int{1, 0}:
		direction = RIGHT
	case [2]int{-1, -1}:
		direction = UP_LEFT
	case [2]int{1, -1}:
		direction = UP_RIGHT
	case [2]int{-1, 1}:
		direction = DOWN_LEFT
	case [2]int{1, 1}:
		direction = DOWN_RIGHT
	default:
		panic(fmt.Sprintf("Invalid direction for (%v, %v)", start, end))
	}
	return direction
}

func ExtendM(grid [][]byte, posX [2]int, posM [2]int) [][2]int {
	var (
		maybeEnd  [][2]int
		direction int
	)
	direction = FindDirection(posX, posM)
	maybeEnd = FindEnd(grid, posM, direction)
	return maybeEnd
}

func ExtendX(grid [][]byte, pos [2]int) [][2]int {
	var (
		ends      [][2]int
		neighbors [][2]int
	)
	ends = make([][2]int, 0)
	neighbors = Neighbors(pos, len(grid), len(grid[0]))
	for _, neighbor := range neighbors {
		char := grid[neighbor[1]][neighbor[0]]
		if char == M {
			ends = append(ends, ExtendM(grid, pos, neighbor)...)
		}
	}
	return ends
}

func CountOccurances(inputs []string) int {
	var (
		count int
		grid  [][]byte
	)
	grid = ParseGrid(inputs)
	for i, row := range grid {
		for j, char := range row {
			if char == X {
				ends := ExtendX(grid, [2]int{j, i})
				count += len(ends)
			}
		}
	}
	return count
}

func ExtendA(grid [][]byte, pos [2]int) bool {
	/*
	 * 0.2
	 * ...
	 * 1.3
	 */
	diagonals := Diagonals(pos, len(grid), len(grid[0]))
	if len(diagonals) != 4 {
		return false
	}
	chars := make([]byte, 4)
	for i, diagonal := range diagonals {
		x, y := diagonal[0], diagonal[1]
		chars[i] = grid[y][x]
	}
	validChars := [4][]byte{
		/*
		 * M.M
		 * ...
		 * S.S
		 */
		{M, S, M, S},
		/*
		 * M.S
		 * ...
		 * M.S
		 */
		{M, M, S, S},
		/*
		 * S.M
		 * ...
		 * S.M
		 */
		{S, S, M, M},
		/*
		 * S.S
		 * ...
		 * M.M
		 */
		{S, M, S, M},
	}
	for _, valid := range validChars {
		if string(valid) == string(chars) {
			return true
		}
	}
	return false
}

func CountOccurancesX(inputs []string) int {
	var (
		count int
		grid  [][]byte
	)
	grid = ParseGrid(inputs)
	for i, row := range grid {
		for j, char := range row {
			if char == A {
				if ExtendA(grid, [2]int{j, i}) {
					count++
				}
			}
		}
	}
	return count
}
