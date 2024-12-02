package day8

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	left     = 'L'
	right    = 'R'
	start    = "AAA"
	end      = "ZZZ"
	maxInt   = int(^uint(0) >> 1)
	maxSteps = 1_000_000
)

var (
	directionRegex = regexp.MustCompile(`^(L|R)+$`)
	nodeRegex      = regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)
	startRegex     = regexp.MustCompile(`^[A-Z0-9]{2}A$`)
	endRegex       = regexp.MustCompile(`^[A-Z0-9]{2}Z$`)
)

type Directions string

func NewDirections(line string) (Directions, error) {
	line = strings.TrimRight(line, "\n")
	if !directionRegex.MatchString(line) {
		return "", fmt.Errorf("invalid directions")
	}
	return Directions(line), nil
}

type Node struct {
	Id    string
	Left  string
	Right string
}

func NewNode(line string) (Node, error) {
	line = strings.TrimRight(line, "\n")
	match := nodeRegex.FindStringSubmatch(line)
	if match == nil {
		return Node{}, fmt.Errorf("invalid node")
	}
	return Node{
		Id:    match[1],
		Left:  match[2],
		Right: match[3],
	}, nil
}

type Graph struct {
	Nodes  map[string]Node
	Starts map[string]bool
	Ends   map[string]bool
}

func NewGraph(lines []string) Graph {
	graph := Graph{make(map[string]Node), make(map[string]bool), make(map[string]bool)}
	for _, line := range lines {
		node, err := NewNode(line)
		if err != nil {
			continue
		}
		graph.Nodes[node.Id] = node
		if startRegex.MatchString(node.Id) {
			graph.Starts[node.Id] = true
		}
		if endRegex.MatchString(node.Id) {
			graph.Ends[node.Id] = true
		}
	}
	return graph
}

func (g Graph) Equals(other Graph) bool {
	if len(g.Nodes) != len(other.Nodes) {
		return false
	}
	for k, v := range g.Nodes {
		if other.Nodes[k] != v {
			return false
		}
	}
	return true
}

type Cycle struct {
	Start  int
	Length int
}

var NonCycle = Cycle{-1, -1}

type Set map[int]bool

func (s Set) Intersection(other Set) Set {
	result := make(Set)
	for k := range s {
		if other[k] {
			result[k] = true
		}
	}
	return result
}

type NodeIndexPair struct {
	Node           Node
	DirectionIndex int
}

type Ghost struct {
	Path           []Node
	PathMap        map[NodeIndexPair]int
	Directions     *Directions
	Graph          *Graph
	Start          string
	EndIndices     []int
	EndIndexSet    Set
	Node           Node
	DirectionIndex int
	PathIndex      int
	Cycle          Cycle
}

func NewGhost(directions *Directions, graph *Graph, start string) Ghost {
	return Ghost{
		Path:           []Node{graph.Nodes[start]},
		PathMap:        map[NodeIndexPair]int{{graph.Nodes[start], 0}: 0},
		Graph:          graph,
		Directions:     directions,
		Start:          start,
		EndIndices:     []int{},
		EndIndexSet:    Set{},
		Node:           graph.Nodes[start],
		DirectionIndex: 0,
		PathIndex:      0,
		Cycle:          NonCycle,
	}
}
func (g *Ghost) Step() {
	if g.Cycle != NonCycle {
		g.DirectionIndex = (g.DirectionIndex + 1) % len(*g.Directions)
		g.PathIndex = g.Cycle.Start + (g.PathIndex+1-g.Cycle.Start)%g.Cycle.Length
		g.Node = g.Path[g.PathIndex]
		return
	}
	direction := (*g.Directions)[g.DirectionIndex]
	if direction == left {
		g.Node = g.Graph.Nodes[g.Node.Left]
	} else {
		g.Node = g.Graph.Nodes[g.Node.Right]
	}
	if g.Graph.Ends[g.Node.Id] {
		g.EndIndices = append(g.EndIndices, len(g.Path))
		g.EndIndexSet[len(g.Path)] = true
	}
	prev, ok := g.PathMap[NodeIndexPair{g.Node, g.DirectionIndex}]
	if ok && len(g.EndIndices) != 0 {
		g.Cycle = Cycle{prev, len(g.Path) - prev}
	} else {
		g.PathMap[NodeIndexPair{g.Node, g.DirectionIndex}] = len(g.Path)
		g.Path = append(g.Path, g.Node)
	}
	g.DirectionIndex = (g.DirectionIndex + 1) % len(*g.Directions)
	g.PathIndex++
}

func (g *Ghost) IsEndStep(step int) bool {
	return g.Cycle.Start+(step-g.Cycle.Start)%g.Cycle.Length == g.EndIndices[0]
}

func (g *Ghost) GenerateEndStep() {
	endStep := g.EndIndices[len(g.EndIndices)-1] + g.Cycle.Length
	g.EndIndices = append(g.EndIndices, endStep)
	g.EndIndexSet[endStep] = true
}

func GCDEuclid(a, b int) int {
	if a == 0 {
		return b
	}
	return GCDEuclid(b%a, a)
}

func LCM(a, b int) int {
	return a * b / GCDEuclid(a, b)
}

func CountSteps(lines []string) int {
	directions, err := NewDirections(lines[0])
	lenDirections := len(directions)
	if err != nil {
		return -1
	}
	graph := NewGraph(lines[1:])
	node := graph.Nodes[start]
	index := 0
	steps := 0
	for {
		direction := directions[index]
		if direction == left {
			node = graph.Nodes[node.Left]
		} else {
			node = graph.Nodes[node.Right]
		}
		steps++
		if node.Id == end {
			break
		}
		index = (index + 1) % lenDirections
	}
	return steps
}

func CountParallelSteps(lines []string) int {
	log.Printf("Running for at most %d steps\n", maxSteps)
	directions, err := NewDirections(lines[0])
	lenDirections := len(directions)
	if err != nil {
		return -1
	}
	graph := NewGraph(lines[1:])
	numGhosts := len(graph.Starts)
	ghosts := make([]Ghost, 0, numGhosts)

	for start := range graph.Starts {
		ghost := NewGhost(&directions, &graph, start)
		ghosts = append(ghosts, ghost)
	}
	index := 0
	steps := 0
	for {
		for i := 0; i < numGhosts; i++ {
			ghosts[i].Step()
		}
		if steps == maxSteps {
			return -1
		}
		steps++
		done := true
		for _, ghost := range ghosts {
			if !graph.Ends[ghost.Node.Id] {
				done = false
				break
			}
		}
		if done {
			break
		}
		allCycling := true
		for _, ghost := range ghosts {
			if ghost.Cycle == NonCycle {
				allCycling = false
				break
			}
		}
		if allCycling {

			/*
				fmt.Printf("All ghosts are cycling after %d steps\n", steps)
				for n := 0; n < 10; n++ {
					commonEndIndices := ghosts[0].EndIndexSet
					for i := 1; i < numGhosts; i++ {
						commonEndIndices = commonEndIndices.Intersection(ghosts[i].EndIndexSet)
					}
					if len(commonEndIndices) == 0 {
						for i := 0; i < numGhosts; i++ {
							ghosts[i].GenerateEndStep()
						}
					} else {
						endStep := maxInt
						for endIndex := range commonEndIndices {
							if endIndex < endStep {
								endStep = endIndex
							}
						}
						return endStep
					}
				}
				for i := 0; i < numGhosts; i++ {
					ghost := ghosts[i]
					fmt.Printf("%#v, EndIndices: %v\n", ghost.Cycle, ghost.EndIndices)
				}
				return -1
			*/
			lcm := 1
			for i := 0; i < numGhosts; i++ {
				ghost := ghosts[i]
				lcm = LCM(lcm, ghost.Cycle.Length)
				fmt.Printf("%#v, EndIndex: %v\n", ghost.Cycle, ghost.EndIndices)
			}
			return lcm
		}
		index = (index + 1) % lenDirections
	}
	return steps
}
