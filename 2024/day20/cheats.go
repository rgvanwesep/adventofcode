package day20

import (
	"aoc2024/set"
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

type vector struct {
	x, y int
}

func dist(v1, v2 vector) int {
	var x, y int
	if x = v1.x - v2.x; x < 0 {
		x = -x
	}
	if y = v1.y - v2.y; y < 0 {
		y = -y
	}
	return x + y
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
	unvisited := set.NewSet[int]()
	for i := range nnodes {
		dists[i] = maxInt
		unvisited.Add(i)
	}
	dists[startId] = 0

	for unvisited.Len() > 0 {
		minValue := maxInt
		for v := range unvisited.All() {
			if dists[v] <= minValue {
				minValue = dists[v]
				u = v
			}
		}
		unvisited.Remove(u)

		if dists[u] == maxInt {
			continue
		}
		for _, conn := range g.adjacencies[u] {
			if unvisited.Contains(conn.nodeId) {
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
	cost      int
}

func (b bridge) getDistanceThrough() int {
	return min(
		b.endpoints[0].distanceFromStart+b.endpoints[1].distanceToEnd,
		b.endpoints[1].distanceFromStart+b.endpoints[0].distanceToEnd,
	) + b.cost
}

type maze struct {
	grid           *grid
	graph          *graph[vector]
	startId, endId int
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
	nodeIds := map[vector]int{}
	for v, c := range gri.all() {
		if c != wallChar {
			nodeIds[v] = gra.addNode(v)
		}
		if c == startChar {
			startId = nodeIds[v]
		} else if c == endChar {
			endId = nodeIds[v]
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
	}
}

func countCheatsBySavings(inputs []string, maxCost int, threshold int) map[int]int {
	var cheat bridge
	cheatsBySavings := map[int]int{}
	m := parseMaze(inputs)
	startDists, _ := dijkstra(m.graph, m.startId)
	endDists, _ := dijkstra(m.graph, m.endId)
	baseline := startDists[m.endId]
	for i, vi := range m.graph.allNodes() {
		ei := endpoint{
			id:       i,
			position: vi,
		}
		for j, vj := range m.graph.allNodes() {
			if j < i {
				ej := endpoint{
					id:       j,
					position: vj,
				}
				cheat = bridge{
					endpoints: [2]endpoint{ei, ej},
				}
				for k := range cheat.endpoints {
					cheat.endpoints[k].distanceFromStart = startDists[cheat.endpoints[k].id]
					cheat.endpoints[k].distanceToEnd = endDists[cheat.endpoints[k].id]
				}
				cheat.cost = dist(cheat.endpoints[0].position, cheat.endpoints[1].position)
				if cheat.cost <= maxCost {
					savings := baseline - cheat.getDistanceThrough()
					if savings >= threshold {
						cheatsBySavings[savings]++
					}
				}
			}
		}
	}
	return cheatsBySavings
}

func CountCheats(inputs []string, maxCost int, threshold int) int {
	cheatsBySavings := countCheatsBySavings(inputs, maxCost, threshold)
	sum := 0
	for _, count := range cheatsBySavings {
		sum += count
	}
	return sum
}
