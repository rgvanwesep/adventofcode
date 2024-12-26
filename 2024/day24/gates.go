package day24

import (
	"fmt"
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

func Evaluate(inputs []string) int {
	var (
		inputA, inputB, output chan bool
		ok                     bool
	)
	initialValues, gates := parseInputs(inputs)
	channels := map[string]chan bool{}
	for _, gate := range gates {
		if inputA, ok = channels[gate.inputA]; !ok {
			inputA = make(chan bool, len(gates))
			channels[gate.inputA] = inputA
		}
		if inputB, ok = channels[gate.inputB]; !ok {
			inputB = make(chan bool, len(gates))
			channels[gate.inputB] = inputB
		}
		if output, ok = channels[gate.output]; !ok {
			output = make(chan bool, len(gates))
			channels[gate.output] = output
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
	for _, initialValue := range initialValues {
		channels[initialValue.name] <- initialValue.value
	}
	result := 0
	bitCount := 0
	for {
		name := fmt.Sprintf("z%02d", bitCount)
		ch, ok := channels[name]
		if !ok {
			break
		}
		if <-ch {
			result ^= 1 << bitCount
		}
		bitCount++
	}
	for _, ch := range channels {
		close(ch)
	}
	return result
}
