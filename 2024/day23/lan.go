package day23

import (
	"iter"
	"slices"
	"strings"
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

func (s *set[T]) clone() set[T] {
	newSet := newSet[T]()
	for v := range s.values {
		newSet.add(v)
	}
	return newSet
}

func intersection[T comparable](s1, s2 set[T]) set[T] {
	s := newSet[T]()
	for i := range s1.all() {
		if s2.contains(i) {
			s.add(i)
		}
	}
	return s
}

type connection struct {
	nodeId     int
	edgeWeight int
}

type graph[T any] struct {
	nodes        []T
	adjacencies  [][]connection
	neighborSets []set[int]
}

func newGraph[T any]() *graph[T] {
	return &graph[T]{
		nodes:        []T{},
		adjacencies:  [][]connection{},
		neighborSets: []set[int]{},
	}
}

func (g *graph[T]) getNode(id int) T {
	return g.nodes[id]
}

func (g *graph[T]) addNode(n T) int {
	id := len(g.nodes)
	g.nodes = append(g.nodes, n)
	g.adjacencies = append(g.adjacencies, []connection{})
	s := newSet[int]()
	g.neighborSets = append(g.neighborSets, s)
	return id
}

func (g *graph[T]) addEdge(id int, conn connection) {
	g.adjacencies[id] = append(g.adjacencies[id], conn)
	g.neighborSets[id].add(conn.nodeId)
}

func (g *graph[T]) getNeighborSet(id int) set[int] {
	return g.neighborSets[id]
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

func (g *graph[T]) expandClique(ids set[int]) iter.Seq[set[int]] {
	return func(yield func(set[int]) bool) {
		var shared set[int]
		for id := range ids.all() {
			shared = g.getNeighborSet(id)
			break
		}
		for id := range ids.all() {
			neighbors := g.getNeighborSet(id)
			shared = intersection(shared, neighbors)
		}
		if shared.size() == 0 {
			if !yield(ids) {
				return
			}
		}
	Outer:
		for id := range shared.all() {
			newClique := ids.clone()
			newClique.add(id)
			for clique := range g.expandClique(newClique) {
				if !yield(clique) {
					break Outer
				}
			}
		}
	}
}

func (g *graph[T]) allCliques() iter.Seq[set[int]] {
	return func(yield func(set[int]) bool) {
		for id := range g.nodes {
			start := newSet[int]()
			start.add(id)
			cliques := g.expandClique(start)
			for clique := range cliques {
				if !yield(clique) {
					break
				}
			}
		}
	}
}

func parseGraph(inputs []string) *graph[string] {
	g := newGraph[string]()
	nodeIds := map[string]int{}
	for _, input := range inputs {
		names := strings.Split(input, "-")
		i, ok := nodeIds[names[0]]
		if !ok {
			i = g.addNode(names[0])
			nodeIds[names[0]] = i
		}
		j, ok := nodeIds[names[1]]
		if !ok {
			j = g.addNode(names[1])
			nodeIds[names[1]] = j
		}
		g.addEdge(i, connection{j, 1})
		g.addEdge(j, connection{i, 1})
	}
	return g
}

func CountLANs(inputs []string) int {
	g := parseGraph(inputs)
	cliques := newSet[[3]int]()
	for id, node := range g.allNodes() {
		if node[0] == 't' {
			neighbors := g.getNeighborSet(id)
			for nid := range neighbors.all() {
				nextNeighbors := g.getNeighborSet(nid)
				shared := intersection(neighbors, nextNeighbors)
				for nnid := range shared.all() {
					lan := []int{id, nid, nnid}
					slices.Sort(lan)
					cliques.add([3]int(lan))
				}
			}
		}
	}
	return cliques.size()
}

func getCliquePassword(g *graph[string], clique set[int]) string {
	names := []string{}
	for id := range clique.all() {
		names = append(names, g.getNode(id))
	}
	slices.Sort(names)
	return strings.Join(names, ",")
}

func FindPassword(inputs []string) string {
	var maxClique set[int]
	maxCliqueSize := 0
	g := parseGraph(inputs)
	for clique := range g.allCliques() {
		cliqueSize := clique.size()
		if cliqueSize > maxCliqueSize {
			maxClique = clique
			maxCliqueSize = cliqueSize
		}
	}
	return getCliquePassword(g, maxClique)
}
