package day24

import (
	"errors"
	"fmt"
	"log"
	"regexp"
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

func computeResult(wires map[string]*wire, initialValues []initialValue, e *executor) int {
	for _, initialValue := range initialValues {
		wires[initialValue.name].send(initialValue.value)
	}
	e.loop()
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
	return result
}

func Evaluate(inputs []string) int {
	initialValues, gates := parseInputs(inputs)
	wires, e := startGates(gates)
	result := computeResult(wires, initialValues, e)
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

func FindSwapped(inputs []string) string {
	initialValues, gates := parseInputs(inputs)
	nBits := len(initialValues) / 2
	for i := range nBits {
		for cx := range 2 {
			for cy := range 2 {
				x := cx << i
				y := cy << i
				initialValues = generateInitalValues(x, y, nBits)
				wires, e := startGates(gates)
				z := computeResult(wires, initialValues, e)
				if z-(x+y) != 0 {
					log.Printf("Result incorrect\n   %045b\n + %045b\n!= %045b", x, y, z)
				}
			}
		}
	}
	return ""
}
