package day18

import (
	"iter"
	"log"
	"strconv"
	"strings"
)

const (
	emptyChar = '.'
	wallChar  = '#'
	maxUint   = ^uint(0)
	maxInt    = int(maxUint >> 1)
)

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

func (s *set[T]) contains(v T) bool {
	_, ok := s.values[v]
	return ok
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

type vector struct {
	x, y int
}

type grid struct {
	nrows, ncols int
	values       []byte
}

func newGrid(nrows, ncols int, fill byte) *grid {
	values := make([]byte, nrows*ncols)
	for i := range values {
		values[i] = fill
	}
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

type connection struct {
	nodeId     int
	edgeWeight int
}

type graph[T any] struct {
	nodes       []T
	adjacencies [][]connection
}

func newGraph[T any]() *graph[T] {
	return &graph[T]{
		nodes:       []T{},
		adjacencies: [][]connection{},
	}
}

func (g *graph[T]) addNode(n T) int {
	id := len(g.nodes)
	g.nodes = append(g.nodes, n)
	g.adjacencies = append(g.adjacencies, []connection{})
	return id
}

func (g *graph[T]) addEdge(id int, conn connection) {
	g.adjacencies[id] = append(g.adjacencies[id], conn)
}

func (g *graph[T]) allNodes() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, n := range g.nodes {
			if !yield(i, n) {
				break
			}
		}
	}
}

func dijkstra[T any](g *graph[T], startId int) ([]int, [][]int) {
	var u int
	nnodes := len(g.nodes)
	dists := make([]int, len(g.nodes))
	prevs := make([][]int, len(g.nodes))
	unvisited := newSet[int]()
	for i := range nnodes {
		dists[i] = maxInt
		unvisited.add(i)
	}
	dists[startId] = 0

	for unvisited.size() > 0 {
		minValue := maxInt
		for v := range unvisited.all() {
			if dists[v] <= minValue {
				minValue = dists[v]
				u = v
			}
		}
		unvisited.remove(u)

		if dists[u] == maxInt {
			continue
		}
		for _, conn := range g.adjacencies[u] {
			if unvisited.contains(conn.nodeId) {
				d := dists[u] + conn.edgeWeight
				if d < dists[conn.nodeId] {
					dists[conn.nodeId] = d
					prevs[conn.nodeId] = []int{u}
				} else if d == dists[conn.nodeId] {
					prevs[conn.nodeId] = append(prevs[conn.nodeId], u)
				}
			}
		}
	}
	return dists, prevs
}

type maze struct {
	grid           *grid
	graph          *graph[vector]
	startId, endId int
}

func parseGrid(inputs []string, nrows, ncols int, nInputs int) *grid {
	g := newGrid(nrows, ncols, emptyChar)
	for i := range nInputs {
		input := inputs[i]
		split := strings.Split(input, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Panicf("Could not parse x from input %q", input)
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Panicf("Could not parse y from input %q", input)
		}
		g.set(vector{x, y}, wallChar)
	}
	return g
}

func parseMaze(inputs []string, nrows, ncols int, nInputs int) maze {
	gri := parseGrid(inputs, nrows, ncols, nInputs)
	gra := newGraph[vector]()
	nodeIds := map[vector]int{}
	for v, c := range gri.all() {
		if c != wallChar {
			nodeIds[v] = gra.addNode(v)
		}
	}
	for nodeId, node := range gra.allNodes() {
		neighbors := [4]vector{
			{node.x, node.y - 1},
			{node.x, node.y + 1},
			{node.x - 1, node.y},
			{node.x + 1, node.y},
		}
		for _, neighbor := range neighbors {
			if c, ok := gri.get(neighbor); ok && c != wallChar {
				gra.addEdge(nodeId, connection{
					nodeId:     nodeIds[neighbor],
					edgeWeight: 1,
				})
			}
		}
	}
	return maze{
		grid:    gri,
		graph:   gra,
		startId: nodeIds[vector{0, 0}],
		endId:   nodeIds[vector{ncols - 1, nrows - 1}],
	}
}

func CountSteps(inputs []string, nrows, ncols int, nInputs int) int {
	maze := parseMaze(inputs, nrows, ncols, nInputs)
	dists, _ := dijkstra(maze.graph, maze.startId)
	return dists[maze.endId]
}

func FindFinalInput(inputs []string, nrows, ncols int) string {
	var i int
	for i = range inputs {
		log.Printf("i, inputs[i] = %d, %s", i, inputs[i])
		if CountSteps(inputs, nrows, ncols, i+1) == maxInt {
			break
		}
	}
	return inputs[i]
}
