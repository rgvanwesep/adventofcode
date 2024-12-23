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
		states = append(states, getAllStatesEndingIn(b)...)
	}
	return states
}

type robot interface {
	getPostion() byte
	setPosition(byte)
	getNext() robot
	setNext(robot)
	move(byte) bool
	press() bool
}

type directionalRobot struct {
	position byte
	next     robot
}

func (r *directionalRobot) getPostion() byte {
	return r.position
}

func (r *directionalRobot) setPosition(position byte) {
	r.position = position
}

func (r *directionalRobot) getNext() robot {
	return r.next
}

func (r *directionalRobot) setNext(next robot) {
	r.next = next
}

/*
	+---+---+
	| ^ | A |

+---+---+---+
| < | v | > |
+---+---+---+
*/
func (r *directionalRobot) move(direction byte) bool {
	switch [2]byte{r.position, direction} {
	case [2]byte{'A', '^'}:
		return false
	case [2]byte{'A', 'v'}:
		r.position = '>'
		return true
	case [2]byte{'A', '<'}:
		r.position = '^'
		return true
	case [2]byte{'A', '>'}:
		return false
	case [2]byte{'^', '^'}:
		return false
	case [2]byte{'^', 'v'}:
		r.position = 'v'
		return true
	case [2]byte{'^', '<'}:
		return false
	case [2]byte{'^', '>'}:
		r.position = 'A'
		return true
	case [2]byte{'v', '^'}:
		r.position = '^'
		return true
	case [2]byte{'v', 'v'}:
		return false
	case [2]byte{'v', '<'}:
		r.position = '<'
		return true
	case [2]byte{'v', '>'}:
		r.position = '>'
		return true
	case [2]byte{'<', '^'}:
		return false
	case [2]byte{'<', 'v'}:
		return false
	case [2]byte{'<', '<'}:
		return false
	case [2]byte{'<', '>'}:
		r.position = 'v'
		return true
	case [2]byte{'>', '^'}:
		r.position = 'A'
		return true
	case [2]byte{'>', 'v'}:
		return false
	case [2]byte{'>', '<'}:
		r.position = 'v'
		return true
	case [2]byte{'>', '>'}:
		return false
	}
	return false
}

func (r *directionalRobot) press() bool {
	if r.next == nil {
		return true
	}
	if r.position == 'A' {
		return r.getNext().press()
	}
	return r.getNext().move(r.position)
}

var _ robot = &directionalRobot{}

type numericalRobot struct {
	position byte
}

func (r *numericalRobot) getPostion() byte {
	return r.position
}

func (r *numericalRobot) setPosition(position byte) {
	r.position = position
}

func (r *numericalRobot) getNext() robot {
	return nil
}

func (r *numericalRobot) setNext(next robot) {
}

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+

	| 0 | A |
	+---+---+
*/
func (r *numericalRobot) move(direction byte) bool {
	switch [2]byte{r.position, direction} {
	case [2]byte{'A', '^'}:
		r.position = '3'
		return true
	case [2]byte{'A', 'v'}:
		return false
	case [2]byte{'A', '<'}:
		r.position = '0'
		return true
	case [2]byte{'A', '>'}:
		return false
	case [2]byte{'0', '^'}:
		r.position = '2'
		return true
	case [2]byte{'0', 'v'}:
		return false
	case [2]byte{'0', '<'}:
		return false
	case [2]byte{'0', '>'}:
		r.position = 'A'
		return true
	case [2]byte{'1', '^'}:
		r.position = '4'
		return true
	case [2]byte{'1', 'v'}:
		return false
	case [2]byte{'1', '<'}:
		return false
	case [2]byte{'1', '>'}:
		r.position = '2'
		return true
	case [2]byte{'2', '^'}:
		r.position = '5'
		return true
	case [2]byte{'2', 'v'}:
		r.position = '0'
		return true
	case [2]byte{'2', '<'}:
		r.position = '1'
		return true
	case [2]byte{'2', '>'}:
		r.position = '3'
		return true
	case [2]byte{'3', '^'}:
		r.position = '6'
		return true
	case [2]byte{'3', 'v'}:
		r.position = 'A'
		return true
	case [2]byte{'3', '<'}:
		r.position = '2'
		return true
	case [2]byte{'3', '>'}:
		return false
	case [2]byte{'4', '^'}:
		r.position = '7'
		return true
	case [2]byte{'4', 'v'}:
		r.position = '1'
		return true
	case [2]byte{'4', '<'}:
		return false
	case [2]byte{'4', '>'}:
		r.position = '5'
		return true
	case [2]byte{'5', '^'}:
		r.position = '8'
		return true
	case [2]byte{'5', 'v'}:
		r.position = '2'
		return true
	case [2]byte{'5', '<'}:
		r.position = '4'
		return true
	case [2]byte{'5', '>'}:
		r.position = '6'
		return true
	case [2]byte{'6', '^'}:
		r.position = '9'
		return true
	case [2]byte{'6', 'v'}:
		r.position = '3'
		return true
	case [2]byte{'6', '<'}:
		r.position = '5'
		return true
	case [2]byte{'6', '>'}:
		return false
	case [2]byte{'7', '^'}:
		return false
	case [2]byte{'7', 'v'}:
		r.position = '4'
		return true
	case [2]byte{'7', '<'}:
		return false
	case [2]byte{'7', '>'}:
		r.position = '8'
		return true
	case [2]byte{'8', '^'}:
		return false
	case [2]byte{'8', 'v'}:
		r.position = '5'
		return true
	case [2]byte{'8', '<'}:
		r.position = '7'
		return true
	case [2]byte{'8', '>'}:
		r.position = '9'
		return true
	case [2]byte{'9', '^'}:
		return false
	case [2]byte{'9', 'v'}:
		r.position = '6'
		return true
	case [2]byte{'9', '<'}:
		r.position = '8'
		return true
	case [2]byte{'9', '>'}:
		return false
	}
	return false
}

func (r *numericalRobot) press() bool {
	return true
}

var _ robot = &numericalRobot{}

type robotChain struct {
	robots [3]robot
}

func newRobotChain(state string) *robotChain {
	robots := [3]robot{
		new(directionalRobot),
		new(directionalRobot),
		new(numericalRobot),
	}
	for i := range 3 {
		robots[i].setPosition(state[i])
	}
	for i := range 2 {
		robots[i].setNext(robots[i+1])
	}
	return &robotChain{robots}
}

func (r *robotChain) apply(input byte) bool {
	if input == 'A' {
		return r.robots[0].press()
	}
	return r.robots[0].move(input)
}

func (r *robotChain) getState() string {
	return string([]byte{
		r.robots[0].getPostion(),
		r.robots[1].getPostion(),
		r.robots[2].getPostion(),
	})
}

func isValidWithInput(input byte, from, to string) bool {
	robots := newRobotChain(from)
	if ok := robots.apply(input); ok && robots.getState() == to {
		return true
	}
	return false
}

func isValidEdge(from, to string) bool {
	directionalValues := []byte{'^', 'v', '<', '>', 'A'}
	for _, input := range directionalValues {
		if isValidWithInput(input, from, to) {
			return true
		}
	}
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
