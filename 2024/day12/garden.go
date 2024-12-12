package day12

import "log"

var regionId int

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
	newRegionFound := false
	for _, neighbor := range neighbors {
		if plot, ok := g.GetValue(neighbor); ok {
			if plot.regionId == -1 {
				var nextRegion Region
				if plot.plant == m.region.plant {
					nextRegion = m.region
				} else {
					if !newRegionFound {
						newRegionFound = true
						regionId++
						nextRegion = Region{regionId, plot.plant}
					}
				}
				mappers = append(mappers, Mapper{neighbor, nextRegion})
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
			plot := Plot{plant: value, regionId: -1}
			if ok := grid.SetValue(p, plot); !ok {
				log.Panicf("Could not set value for %v", p)
			}
		}
	}
	return grid
}

func SumFencePrice(inputs []string) int {
	grid := ParseGrid(inputs)
	mappers := NewQueue[Mapper]()
	plot, _ := grid.GetValue(Point{0,0})
	mappers.Push(Mapper{position: Point{0, 0}, region: Region{regionId, plot.plant}})
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
	return 0
}
