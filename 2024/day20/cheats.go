package day20

import (
	"iter"
)

const (
	emptyChar = '.'
	wallChar  = '#'
	startChar = 'S'
	endChar   = 'E'
	maxUint   = ^uint(0)
	maxInt    = int(maxUint >> 1)
	threshold = 100
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

type endpoint struct {
	id                int
	position          vector
	distanceFromStart int
	distanceToEnd     int
}

type bridge struct {
	endpoints [2]endpoint
}

func (b bridge) getDistanceThrough() int {
	return min(
		b.endpoints[0].distanceFromStart+b.endpoints[1].distanceToEnd,
		b.endpoints[1].distanceFromStart+b.endpoints[0].distanceToEnd,
	) + 2
}

type maze struct {
	grid           *grid
	graph          *graph[vector]
	startId, endId int
	cheats         []bridge
}

func parseGrid(inputs []string) *grid {
	nrows := len(inputs)
	ncols := len(inputs[0])
	g := newGrid(nrows, ncols)
	for i, row := range inputs {
		for j := range row {
			v := vector{j, i}
			c := row[j]
			g.set(v, c)
		}
	}
	return g
}

func parseMaze(inputs []string) maze {
	var startId, endId int
	gri := parseGrid(inputs)
	gra := newGraph[vector]()
	cheats := []bridge{}
	nodeIds := map[vector]int{}
	for v, c := range gri.all() {
		if c == wallChar {
			v1 := vector{v.x - 1, v.y}
			v2 := vector{v.x + 1, v.y}
			c1, ok1 := gri.get(v1)
			c2, ok2 := gri.get(v2)
			if ok1 && ok2 && c1 != wallChar && c2 != wallChar {
				cheats = append(cheats, bridge{
					endpoints: [2]endpoint{
						{position: v1},
						{position: v2},
					},
				})
			}

			v1 = vector{v.x, v.y - 1}
			v2 = vector{v.x, v.y + 1}
			c1, ok1 = gri.get(v1)
			c2, ok2 = gri.get(v2)
			if ok1 && ok2 && c1 != wallChar && c2 != wallChar {
				cheats = append(cheats, bridge{
					endpoints: [2]endpoint{
						{position: v1},
						{position: v2},
					},
				})
			}
		} else {
			nodeIds[v] = gra.addNode(v)
		}
		if c == startChar {
			startId = nodeIds[v]
		} else if c == endChar {
			endId = nodeIds[v]
		}
	}
	for i := range cheats {
		for j := range cheats[i].endpoints {
			cheats[i].endpoints[j].id = nodeIds[cheats[i].endpoints[j].position]
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
		startId: startId,
		endId:   endId,
		cheats:  cheats,
	}
}

func countCheatsBySavings(inputs []string) map[int]int {
	cheatsBySavings := map[int]int{}
	m := parseMaze(inputs)
	startDists, _ := dijkstra(m.graph, m.startId)
	endDists, _ := dijkstra(m.graph, m.endId)
	baseline := startDists[m.endId]
	for i := range m.cheats {
		for j := range m.cheats[i].endpoints {
			m.cheats[i].endpoints[j].distanceFromStart = startDists[m.cheats[i].endpoints[j].id]
			m.cheats[i].endpoints[j].distanceToEnd = endDists[m.cheats[i].endpoints[j].id]
		}
		savings := baseline - m.cheats[i].getDistanceThrough()
		cheatsBySavings[savings]++
	}
	return cheatsBySavings
}

func CountCheats(inputs []string) int {
	cheatsBySavings := countCheatsBySavings(inputs)
	sum := 0
	for savings, count := range cheatsBySavings {
		if savings >= threshold {
			sum += count
		}
	}
	return sum
}
