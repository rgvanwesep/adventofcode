package day16

import "iter"

const (
	emptyChar = '.'
	wallChar  = '#'
	startChar = 'S'
	endChar   = 'E'
	upChar    = '^'
	downChar  = 'v'
	leftChar  = '<'
	rightChar = '>'
	maxUint   = ^uint(0)
	maxInt    = int(maxUint >> 1)
)

type stack[T any] struct {
	values []T
}

func newStack[T any]() stack[T] {
	return stack[T]{
		values: []T{},
	}
}

func (s *stack[T]) push(v T) {
	s.values = append(s.values, v)
}

func (s *stack[T]) pop() (v T, ok bool) {
	size := len(s.values)
	if size > 0 {
		v = s.values[size-1]
		s.values = s.values[:size-1]
		ok = true
	}
	return
}

func (s *stack[T]) clear() {
	s.values = []T{}
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
			if dists[v] < minValue {
				minValue = dists[v]
				u = v
			}
		}
		unvisited.remove(u)

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

type state struct {
	position    vector
	orientation byte
}

type maze struct {
	grid    *grid
	graph   *graph[state]
	startId int
	endIds  [4]int
}

func parseGrid(inputs []string) *grid {
	nrows := len(inputs)
	ncols := len(inputs[0])
	g := newGrid(nrows, ncols)
	for i, row := range inputs {
		for j := range row {
			v := vector{j, nrows - i - 1}
			c := row[j]
			g.set(v, c)
		}
	}
	return g
}

func parseMaze(inputs []string) maze {
	var (
		startId                                   int
		endIds                                    [4]int
		upState, downState, leftState, rightState state
	)
	gri := parseGrid(inputs)
	gra := newGraph[state]()
	nodeIds := map[state]int{}
	for v, c := range gri.all() {
		if c != wallChar {
			upState = state{v, upChar}
			nodeIds[upState] = gra.addNode(upState)
			downState = state{v, downChar}
			nodeIds[downState] = gra.addNode(downState)
			leftState = state{v, leftChar}
			nodeIds[leftState] = gra.addNode(leftState)
			rightState = state{v, rightChar}
			nodeIds[rightState] = gra.addNode(rightState)
		}
		if c == startChar {
			startId = nodeIds[rightState]
		} else if c == endChar {
			endIds = [4]int{
				nodeIds[upState],
				nodeIds[downState],
				nodeIds[leftState],
				nodeIds[rightState],
			}
		}
	}
	for nodeId, node := range gra.allNodes() {
		switch node.orientation {
		case upChar:
			forward := state{
				position:    vector{node.position.x, node.position.y + 1},
				orientation: upChar,
			}
			if gri.get(forward.position) != wallChar {
				gra.addEdge(nodeId, connection{
					nodeId:     nodeIds[forward],
					edgeWeight: 1,
				})
			}

			turnLeft := state{
				position:    vector{node.position.x, node.position.y},
				orientation: leftChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnLeft],
				edgeWeight: 1000,
			})

			turnRight := state{
				position:    vector{node.position.x, node.position.y},
				orientation: rightChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnRight],
				edgeWeight: 1000,
			})
		case downChar:
			forward := state{
				position:    vector{node.position.x, node.position.y - 1},
				orientation: downChar,
			}
			if gri.get(forward.position) != wallChar {
				gra.addEdge(nodeId, connection{
					nodeId:     nodeIds[forward],
					edgeWeight: 1,
				})
			}

			turnLeft := state{
				position:    vector{node.position.x, node.position.y},
				orientation: rightChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnLeft],
				edgeWeight: 1000,
			})

			turnRight := state{
				position:    vector{node.position.x, node.position.y},
				orientation: leftChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnRight],
				edgeWeight: 1000,
			})
		case leftChar:
			forward := state{
				position:    vector{node.position.x - 1, node.position.y},
				orientation: leftChar,
			}
			if gri.get(forward.position) != wallChar {
				gra.addEdge(nodeId, connection{
					nodeId:     nodeIds[forward],
					edgeWeight: 1,
				})
			}

			turnLeft := state{
				position:    vector{node.position.x, node.position.y},
				orientation: downChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnLeft],
				edgeWeight: 1000,
			})

			turnRight := state{
				position:    vector{node.position.x, node.position.y},
				orientation: upChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnRight],
				edgeWeight: 1000,
			})
		case rightChar:
			forward := state{
				position:    vector{node.position.x + 1, node.position.y},
				orientation: rightChar,
			}
			if gri.get(forward.position) != wallChar {
				gra.addEdge(nodeId, connection{
					nodeId:     nodeIds[forward],
					edgeWeight: 1,
				})
			}

			turnLeft := state{
				position:    vector{node.position.x, node.position.y},
				orientation: upChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnLeft],
				edgeWeight: 1000,
			})

			turnRight := state{
				position:    vector{node.position.x, node.position.y},
				orientation: downChar,
			}
			gra.addEdge(nodeId, connection{
				nodeId:     nodeIds[turnRight],
				edgeWeight: 1000,
			})
		}
	}
	return maze{
		grid:    gri,
		graph:   gra,
		startId: startId,
		endIds:  endIds,
	}
}

func MinScore(inputs []string) int {
	maze := parseMaze(inputs)
	dists, _ := dijkstra(maze.graph, maze.startId)
	return min(
		dists[maze.endIds[0]],
		dists[maze.endIds[1]],
		dists[maze.endIds[2]],
		dists[maze.endIds[3]],
	)
}

func CountTiles(inputs []string) int {
	maze := parseMaze(inputs)
	dists, prevs := dijkstra(maze.graph, maze.startId)
	endIdIndex := 0
	minDist := dists[maze.endIds[endIdIndex]]
	for j := range maze.endIds[1:] {
		if dists[maze.endIds[j]] < minDist {
			minDist = dists[maze.endIds[j]]
			endIdIndex = j
		}
	}
	ids := newStack[int]()
	ids.push(maze.endIds[endIdIndex])
	positions := newSet[vector]()
	for {
		if id, ok := ids.pop(); ok {
			node := maze.graph.nodes[id]
			positions.add(node.position)
			for _, prevId := range prevs[id] {
				ids.push(prevId)
			}
		} else {
			break
		}
	}
	return positions.size()
}
