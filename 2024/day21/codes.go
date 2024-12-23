package day21

import (
	"iter"
	"log"
	"strconv"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
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

func getNumericPart(input string) int {
	inputBytes := []byte(input)
	numericString := string(inputBytes[:len(input)-1])
	if i, err := strconv.Atoi(numericString); err == nil {
		return i
	}
	log.Panicf("Could not parse numeric part from input %q", input)
	return 0
}

func getAllStatesEndingIn(b byte) []string {
	states := []string{}
	directionalValues := []byte{'^', 'v', '<', '>', 'A'}
	for _, directionalA := range directionalValues {
		for _, directionalB := range directionalValues {
			states = append(states, string([]byte{directionalA, directionalB, b}))
		}
	}
	return states
}

func getAllStates() []string {
	states := []string{}
	numericValues := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'}
	for _, b := range numericValues {
		for _, state := range getAllStatesEndingIn(b) {
			states = append(states, state)
		}
	}
	return states
}

func isValidEdge(from string, to string) bool {
	return false
}

func getShortestSequenceLength(input string) int {
	g := newGraph[string]()
	nodeIds := map[string]int{}
	for _, state := range getAllStates() {
		nodeIds[state] = g.addNode(state)
	}
	for from, i := range nodeIds {
		for to, j := range nodeIds {
			if isValidEdge(from, to) {
				g.addEdge(i, connection{j, 1})
			}
		}
	}
	totalDist := 0
	start := "AAA"
	for i := range input {
		end := string([]byte{'A', 'A', input[i]})
		dists, _ := dijkstra(g, nodeIds[start])
		totalDist += dists[nodeIds[end]] + 1
		start = end
	}
	return totalDist
}

func CalcComplexity(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += getNumericPart(input) * getShortestSequenceLength(input)
	}
	return sum
}
