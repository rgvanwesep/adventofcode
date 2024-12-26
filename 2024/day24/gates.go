package day24

import (
	"fmt"
	"log"
	"regexp"
)

const (
	initialValuePattern = `^((x|y)[0-9]{2}): (0|1)$`
	gatePattern         = `^([a-z][0-9a-z]{2}) ((AND)|(OR)|(XOR)) ([a-z][0-9a-z]{2}) -> ([a-z][0-9a-z]{2})$`
)

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

func and(inputA, inputB chan bool, output chan<- bool) {
	var a, b, ok bool
	if a, ok = <-inputA; !ok {
		return
	}
	inputA <- a
	if b, ok = <-inputB; !ok {
		return
	}
	inputB <- b
	output <- a && b
}

func or(inputA, inputB chan bool, output chan<- bool) {
	var a, b, ok bool
	if a, ok = <-inputA; !ok {
		return
	}
	inputA <- a
	if b, ok = <-inputB; !ok {
		return
	}
	inputB <- b
	output <- a || b
}

func xor(inputA, inputB chan bool, output chan<- bool) {
	var a, b, ok bool
	if a, ok = <-inputA; !ok {
		return
	}
	inputA <- a
	if b, ok = <-inputB; !ok {
		return
	}
	inputB <- b
	output <- a != b
}

func startGates(gates []gate) map[string]chan bool {
	var (
		inputA, inputB, output chan bool
		ok                     bool
	)
	wires := map[string]chan bool{}
	for _, gate := range gates {
		if inputA, ok = wires[gate.inputA]; !ok {
			inputA = make(chan bool, len(gates))
			wires[gate.inputA] = inputA
		}
		if inputB, ok = wires[gate.inputB]; !ok {
			inputB = make(chan bool, len(gates))
			wires[gate.inputB] = inputB
		}
		if output, ok = wires[gate.output]; !ok {
			output = make(chan bool, len(gates))
			wires[gate.output] = output
		}
		switch gate.operation {
		case "AND":
			go and(inputA, inputB, output)
		case "OR":
			go or(inputA, inputB, output)
		case "XOR":
			go xor(inputA, inputB, output)
		}
	}
	return wires
}

func computeResult(wires map[string]chan bool, initialValues []initialValue) int {
	for _, initialValue := range initialValues {
		wires[initialValue.name] <- initialValue.value
	}
	result := 0
	bitCount := 0
	for {
		name := fmt.Sprintf("z%02d", bitCount)
		wire, ok := wires[name]
		if !ok {
			break
		}
		if <-wire {
			result ^= 1 << bitCount
		}
		bitCount++
	}
	return result
}

func Evaluate(inputs []string) int {
	initialValues, gates := parseInputs(inputs)
	wires := startGates(gates)
	result := computeResult(wires, initialValues)
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
				wires := startGates(gates)
				z := computeResult(wires, initialValues)
				if z-(x+y) != 0 {
					log.Printf("Result incorrect\n   %045b\n + %045b\n!= %045b", x, y, z)
				}
			}
		}
	}
	return ""
}
