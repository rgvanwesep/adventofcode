package day12

import (
	"iter"
	"log"
)

type Point struct {
	x, y int
}

func (p Point) IsNeighbor(other Point) bool {
	sqDist := (p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y)
	return sqDist == 1
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

func (g *Grid[T]) All() iter.Seq2[Point, T] {
	iToP := func(i int) Point {
		row := i / g.ncols
		return Point{i % g.ncols, g.nrows - row - 1}
	}
	return func(yield func(Point, T) bool) {
		for i := 0; i < len(g.values) && yield(iToP(i), g.values[i]); i++ {
		}
	}
}

type Queue[T any] struct {
	values []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		values: []T{},
	}
}

func (q *Queue[T]) Pop() (value T, ok bool) {
	if len(q.values) > 0 {
		value = q.values[0]
		q.values = q.values[1:]
		ok = true
	}
	return
}

func (q *Queue[T]) Push(value T) {
	q.values = append(q.values, value)
}

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		values: []T{},
	}
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	size := len(s.values)
	if size > 0 {
		value = s.values[size-1]
		s.values = s.values[:size-1]
		ok = true
	}
	return
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

type Plot struct {
	plant    byte
	regionId int
}

type Region struct {
	id    int
	plant byte
}

type Mapper struct {
	position Point
	region   Region
}

func (m Mapper) Step(g *Grid[Plot]) []Mapper {
	g.SetValue(m.position, Plot{m.region.plant, m.region.id})
	neighbors := []Point{
		{m.position.x - 1, m.position.y},
		{m.position.x + 1, m.position.y},
		{m.position.x, m.position.y - 1},
		{m.position.x, m.position.y + 1},
	}
	mappers := []Mapper{}
	for _, neighbor := range neighbors {
		if plot, ok := g.GetValue(neighbor); ok {
			if plot.regionId == 0 && plot.plant == m.region.plant {
				mappers = append(mappers, Mapper{neighbor, m.region})
			}
		}
	}
	return mappers
}

func ParseGrid(inputs []string) *Grid[Plot] {
	nrows, ncols := len(inputs), len(inputs[0])
	grid := NewGrid[Plot](nrows, ncols)
	for row, input := range inputs {
		for col := range input {
			p := Point{col, nrows - row - 1}
			value := input[col]
			plot := Plot{plant: value}
			if ok := grid.SetValue(p, plot); !ok {
				log.Panicf("Could not set value for %v", p)
			}
		}
	}
	return grid
}

func AssignRegions(grid *Grid[Plot]) int {
	regionId := 1
	for p, plot := range grid.All() {
		if plot.regionId == 0 {
			mappers := NewStack[Mapper]()
			mappers.Push(Mapper{position: p, region: Region{regionId, plot.plant}})
			for {
				if mapper, ok := mappers.Pop(); ok {
					newMappers := mapper.Step(grid)
					for _, m := range newMappers {
						mappers.Push(m)
					}
				} else {
					break
				}
			}
			regionId++
		}
	}
	return regionId - 1
}

func GetRegionPoints(grid *Grid[Plot], regionId int) []Point {
	points := []Point{}
	for p, plot := range grid.All() {
		if plot.regionId == regionId {
			points = append(points, p)
		}
	}
	return points
}

func GetPerimeter(points []Point) int {
	perimeter := 4 * len(points)
	for i, p1 := range points {
		for _, p2 := range points[:i] {
			if p1.IsNeighbor(p2) {
				perimeter -= 2
			}
		}
	}
	return perimeter
}

func SumFencePrice(inputs []string) int {
	sum := 0
	grid := ParseGrid(inputs)
	log.Printf("Grid has %d points", grid.nrows*grid.ncols)
	nRegions := AssignRegions(grid)
	log.Printf("Grid has %d regions", nRegions)
	for i := 1; i <= nRegions; i++ {
		points := GetRegionPoints(grid, i)
		area := len(points)
		perimeter := GetPerimeter(points)
		sum += area * perimeter
	}
	return sum
}

type Graph struct {
	adjList [][]int
}

func PointsToGraph(points []Point) Graph {
	adjList := make([][]int, len(points))
	for i1, p1 := range points {
		adjList[i1] = make([]int, 0, 4)
		for i2, p2 := range points[:i1] {
			if p1.IsNeighbor(p2) {
				adjList[i1] = append(adjList[i1], i2)
			}
		}
	}
	return Graph{adjList}
}

func (g Graph) GetNeighbors(i int) []int {
	return g.adjList[i]
}

func IsSquare(p0, p1, p2, p3 Point) bool {
	return false
}

func GetNumSides(points []Point, graph Graph) int {
	nCorners := 0
	for i, point := range points {
		neighbors := graph.GetNeighbors(i)
		nextNeighbors := []int{}
		for _, neighbor := range neighbors {
			for _, nextNeighbor := range graph.GetNeighbors(neighbor) {
				addNextNeighbor := nextNeighbor != i
				for _, n := range neighbors {
					if nextNeighbor == n {
						addNextNeighbor = false
						break
					}
				}
				if addNextNeighbor {
					nextNeighbors = append(nextNeighbors, nextNeighbor)
				}
			}
		}
		switch len(neighbors) {
		case 0:
			nCorners += 4
		case 1:
			nCorners += 2
		case 2:
			for _, nextNeighbor := range nextNeighbors {
				if IsSquare(point, points[neighbors[0]], points[neighbors[1]], points[nextNeighbor]) {
					nCorners += 1
				} else {
					nCorners += 2
				}
			}
		}
	}
	return nCorners
}

func SumFencePriceDiscount(inputs []string) int {
	sum := 0
	grid := ParseGrid(inputs)
	log.Printf("Grid has %d points", grid.nrows*grid.ncols)
	nRegions := AssignRegions(grid)
	log.Printf("Grid has %d regions", nRegions)
	for i := 1; i <= nRegions; i++ {
		points := GetRegionPoints(grid, i)
		area := len(points)
		graph := PointsToGraph(points)
		nSides := GetNumSides(points, graph)
		sum += area * nSides
	}
	return sum
}
