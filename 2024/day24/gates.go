package day24

import (
	"errors"
	"fmt"
	"iter"
	"log"
	//"maps"
	"regexp"
	//"slices"
)

const (
	initialValuePattern = `^((x|y)[0-9]{2}): (0|1)$`
	gatePattern         = `^([a-z][0-9a-z]{2}) ((AND)|(OR)|(XOR)) ([a-z][0-9a-z]{2}) -> ([a-z][0-9a-z]{2})$`
)

type wire struct {
	value []bool
}

func newWire() *wire {
	return &wire{
		value: make([]bool, 0, 1),
	}
}

func (w *wire) send(value bool) {
	if len(w.value) == 0 {
		w.value = append(w.value, value)
	} else {
		w.value[0] = value
	}
}

func (w *wire) recv() (value bool, ok bool) {
	if len(w.value) == 1 {
		value = w.value[0]
		ok = true
	}
	return
}

func (w *wire) hasValue() bool {
	return len(w.value) == 1
}

type executor struct {
	gates []struct {
		gate                   func(*wire, *wire, *wire)
		inputA, inputB, output *wire
	}
}

func newExecutor() *executor {
	return new(executor)
}

func (e *executor) exec(f func(*wire, *wire, *wire), inputA, inputB, output *wire) {
	e.gates = append(e.gates, struct {
		gate   func(*wire, *wire, *wire)
		inputA *wire
		inputB *wire
		output *wire
	}{
		f,
		inputA,
		inputB,
		output,
	})
}

func (e *executor) loop() error {
	remaining := len(e.gates)
	for {
		remainingStart := remaining
		for _, gate := range e.gates {
			if !gate.output.hasValue() && gate.inputA.hasValue() && gate.inputB.hasValue() {
				gate.gate(gate.inputA, gate.inputB, gate.output)
				remaining--
			}
		}
		if remaining == 0 {
			break
		}
		if remaining == remainingStart {
			return errors.New("deadlock")
		}
	}
	return nil
}

type initialValue struct {
	name  string
	value bool
}

func (v initialValue) String() string {
	if v.value {
		return fmt.Sprintf("%s: 1", v.name)
	}
	return fmt.Sprintf("%s: 0", v.name)
}

type gate struct {
	operation              string
	inputA, inputB, output string
}

func parseInputs(inputs []string) ([]initialValue, []gate) {
	inInitialValues := true
	initialValues := []initialValue{}
	gates := []gate{}
	matcher := regexp.MustCompile(initialValuePattern)
	for _, input := range inputs {
		if input == "" {
			inInitialValues = false
			matcher = regexp.MustCompile(gatePattern)
			continue
		}
		match := matcher.FindStringSubmatch(input)
		if inInitialValues {
			name := match[1]
			switch digit := match[3]; digit {
			case "0":
				initialValues = append(initialValues, initialValue{name, false})
			case "1":
				initialValues = append(initialValues, initialValue{name, true})
			}
		} else {
			inputA := match[1]
			operation := match[2]
			inputB := match[6]
			output := match[7]
			gates = append(gates, gate{operation, inputA, inputB, output})
		}
	}
	return initialValues, gates
}

func and(inputA, inputB, output *wire) {
	var a, b, ok bool
	if a, ok = inputA.recv(); !ok {
		return
	}
	inputA.send(a)
	if b, ok = inputB.recv(); !ok {
		return
	}
	inputB.send(b)
	output.send(a && b)
}

func or(inputA, inputB, output *wire) {
	var a, b, ok bool
	if a, ok = inputA.recv(); !ok {
		return
	}
	inputA.send(a)
	if b, ok = inputB.recv(); !ok {
		return
	}
	inputB.send(b)
	output.send(a || b)
}

func xor(inputA, inputB, output *wire) {
	var a, b, ok bool
	if a, ok = inputA.recv(); !ok {
		return
	}
	inputA.send(a)
	if b, ok = inputB.recv(); !ok {
		return
	}
	inputB.send(b)
	output.send(a != b)
}

func startGates(gates []gate) (map[string]*wire, *executor) {
	var (
		inputA, inputB, output *wire
		ok                     bool
	)
	wires := map[string]*wire{}
	e := newExecutor()
	for _, gate := range gates {
		if inputA, ok = wires[gate.inputA]; !ok {
			inputA = newWire()
			wires[gate.inputA] = inputA
		}
		if inputB, ok = wires[gate.inputB]; !ok {
			inputB = newWire()
			wires[gate.inputB] = inputB
		}
		if output, ok = wires[gate.output]; !ok {
			output = newWire()
			wires[gate.output] = output
		}
		switch gate.operation {
		case "AND":
			e.exec(and, inputA, inputB, output)
		case "OR":
			e.exec(or, inputA, inputB, output)
		case "XOR":
			e.exec(xor, inputA, inputB, output)
		}
	}
	return wires, e
}

func computeResult(wires map[string]*wire, initialValues []initialValue, e *executor) (int, error) {
	for _, initialValue := range initialValues {
		wires[initialValue.name].send(initialValue.value)
	}
	err := e.loop()
	if err != nil {
		return 0, err
	}
	result := 0
	bitCount := 0
	for {
		name := fmt.Sprintf("z%02d", bitCount)
		wire, ok := wires[name]
		if !ok {
			break
		}
		if value, _ := wire.recv(); value {
			result ^= 1 << bitCount
		}
		bitCount++
	}
	return result, nil
}

func Evaluate(inputs []string) int {
	initialValues, gates := parseInputs(inputs)
	wires, e := startGates(gates)
	result, _ := computeResult(wires, initialValues, e)
	return result
}

func generateInitalValues(x, y, nBits int) []initialValue {
	initialValues := []initialValue{}
	for i := range nBits {
		initialValues = append(initialValues,
			initialValue{
				name:  fmt.Sprintf("x%02d", i),
				value: x&(1<<i) != 0,
			},
			initialValue{
				name:  fmt.Sprintf("y%02d", i),
				value: y&(1<<i) != 0,
			},
		)
	}
	return initialValues
}

type node struct {
	id             string
	childA, childB *node
}

type tree struct {
	nodes      map[string]*node
	edgeValues map[[2]string]string
}

func newTree() *tree {
	return &tree{
		nodes:      map[string]*node{},
		edgeValues: map[[2]string]string{},
	}
}

func (t *tree) addNode(id string) {
	t.nodes[id] = &node{id: id}
}

func (t *tree) addEdge(parentId, childId, value string) bool {
	if parent, ok := t.nodes[parentId]; ok {
		if child, ok := t.nodes[childId]; ok {
			if parent.childA == nil {
				parent.childA = child
				t.edgeValues[[2]string{parentId, childId}] = value
				return true
			}
			if parent.childB == nil {
				parent.childB = child
				t.edgeValues[[2]string{parentId, childId}] = value
				return true
			}
		}
	}
	return false
}

func (t *tree) connectedNodes(rootId string) iter.Seq2[string, *node] {
	return func(yield func(string, *node) bool) {
		nodeIds := []string{rootId}
		size := 1
		for size > 0 {
			size--
			nodeId := nodeIds[size]
			nodeIds = nodeIds[:size]
			node := t.nodes[nodeId]
			if yield(nodeId, node) {
				if node.childA != nil && node.childB != nil {
					nodeIds = append(nodeIds, node.childA.id, node.childB.id)
					size += 2
				} else if node.childA != nil {
					nodeIds = append(nodeIds, node.childA.id)
					size += 1
				} else if node.childB != nil {
					nodeIds = append(nodeIds, node.childB.id)
					size += 1
				}
			} else {
				break
			}
		}
	}
}

func formTree(gates []gate) *tree {
	t := newTree()
	for _, gate := range gates {
		if _, ok := t.nodes[gate.inputA]; !ok {
			t.addNode(gate.inputA)
		}
		if _, ok := t.nodes[gate.inputB]; !ok {
			t.addNode(gate.inputB)
		}
		if _, ok := t.nodes[gate.output]; !ok {
			t.addNode(gate.output)
		}
	}
	for _, gate := range gates {
		t.addEdge(gate.output, gate.inputA, gate.operation)
		t.addEdge(gate.output, gate.inputB, gate.operation)
	}
	return t
}

func getIncorrectOutputs(x, y, z int) []string {
	outputs := []string{}
	mismatched := z ^ (x + y)
	for i := 0; mismatched != 0; i++ {
		if mismatched&1 == 1 {
			outputs = append(outputs, fmt.Sprintf("z%02d", i))
		}
		mismatched >>= 1
	}
	return outputs
}

func swapWires(gates []gate, swaps map[string]string) []gate {
	if len(swaps) == 0 {
		return gates
	}
	swappedGates := make([]gate, len(gates))
	for nameA, nameB := range swaps {
		for i, gate := range gates {
			if swappedGates[i].operation == "" {
				swappedGates[i] = gate
			}
			if gate.inputA == nameA {
				swappedGates[i].inputA = nameB
			}
			if gate.inputA == nameB {
				swappedGates[i].inputA = nameA
			}
			if gate.inputB == nameA {
				swappedGates[i].inputB = nameB
			}
			if gate.inputB == nameB {
				swappedGates[i].inputB = nameA
			}
			if gate.output == nameA {
				swappedGates[i].output = nameB
			}
			if gate.output == nameB {
				swappedGates[i].output = nameA
			}
		}
	}
	return swappedGates
}

func swapOutputs(gates []gate, swaps map[string]string) []gate {
	if len(swaps) == 0 {
		return gates
	}
	swappedGates := make([]gate, len(gates))
	for nameA, nameB := range swaps {
		for i, gate := range gates {
			if swappedGates[i].operation == "" {
				swappedGates[i] = gate
			}
			if gate.output == nameA {
				swappedGates[i].output = nameB
			}
			if gate.output == nameB {
				swappedGates[i].output = nameA
			}
		}
	}
	return swappedGates
}

func FindSwapped(inputs []string) string {
	_, gates := parseInputs(inputs)
	wireCounts := map[string]int{}
	for _, gate := range gates {
		wireCounts[gate.inputA]++
		wireCounts[gate.inputB]++
		wireCounts[gate.output]++
	}
	outputWireCounts := map[string]int{}
	for _, gate := range gates {
		outputWireCounts[gate.output]++
	}
	/*
		for name, count := range wireCounts {
			if count == 1 && name[0] != 'z' {
				log.Printf("Anomolous wire: %s", name)
			}
		}
	*/
	log.Printf("There are %d gates, %d wires and %d output wires",
		len(gates), len(wireCounts), len(outputWireCounts),
	)
	//nBits := len(initialValues) / 2
	nBits := 14
	/*
		allValid := true
		validatedWires := map[string]bool{}
	*/
	wireNames := []string{}
	for name := range outputWireCounts {
		wireNames = append(wireNames, name)
	}
iSwapLoop:
	for iSwap := range wireNames {
	jSwapLoop:
		for jSwap := range wireNames {
			nameA := wireNames[iSwap]
			nameB := wireNames[jSwap]
			swappedGates := swapOutputs(gates, map[string]string{"z12": nameA, "z21": nameB})
			allCorrect := true
			//dependencies := formTree(swappedGates)
			for i := range nBits {
				for cx := range 2 {
					for cy := range 2 {
						x := cx << i
						y := cy << i
						initialValues := generateInitalValues(x, y, nBits)
						wires, e := startGates(swappedGates)
						z, err := computeResult(wires, initialValues, e)
						if err != nil {
							continue jSwapLoop
						}
						if z-(x+y) != 0 {
							allCorrect = false
							//allValid = false
							log.Printf("Result incorrect\n   %045b\n + %045b\n!= %045b", x, y, z)
							incorrectOutputs := getIncorrectOutputs(x, y, z)
							log.Printf("Incorrect outputs: %v", incorrectOutputs)
							/*
								dependentWires := map[string]bool{}
								for _, incorrectOutput := range incorrectOutputs {
									//log.Printf("Gates for output %s:", incorrectOutput)
									for wireName, _ := range dependencies.connectedNodes(incorrectOutput) {

											if i < 14 && node.childA != nil {
												op := dependencies.edgeValues[[2]string{wireName, node.childA.id}]
												log.Printf("%s %s %s -> %s", node.childA.id, op, node.childB.id, wireName)
											}

										dependentWires[wireName] = true
									}
								}
								problemWires := []string{}
								for wireName := range maps.Keys(dependentWires) {
									if !validatedWires[wireName] {
										problemWires = append(problemWires, wireName)
									}
								}
								slices.Sort(problemWires)
								log.Printf("Problem wires: %q", problemWires)
							*/
						}
					}
				}
				/*
					if allValid {
						for wireName, _ := range dependencies.connectedNodes(fmt.Sprintf("z%02d", i)) {
							validatedWires[wireName] = true
						}
					}
				*/
			}
			if allCorrect {
				log.Printf("All correct for swap %s <-> %s", nameA, nameB)
				break iSwapLoop
			}
		}
	}
	return ""
}
