package day17

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

const (
	valA = 4
	valB = 5
	valC = 6
)

const (
	regA = 0
	regB = 1
	regC = 2
)

const (
	regPattern  = `^Register (A|B|C): (\d+)$`
	progPattern = `^Program: ([0-7,]+)$`
)

type instruction struct {
	opcode  uint8
	operand uint8
}

type emulator[T ~int | ~int32 | ~int64] struct {
	registers [3]T
	program   []instruction
	instPtr   int
}

func (e *emulator[T]) execute() string {
	outputs := []string{}
	for e.instPtr < len(e.program) {
		switch inst := e.program[e.instPtr]; inst.opcode {
		case adv:
			switch inst.operand {
			case valA:
				e.registers[regA] >>= e.registers[regA]
			case valB:
				e.registers[regA] >>= e.registers[regB]
			case valC:
				e.registers[regA] >>= e.registers[regC]
			default:
				e.registers[regA] >>= inst.operand
			}
			e.instPtr++
		case bxl:
			e.registers[regB] ^= T(inst.operand)
			e.instPtr++
		case bst:
			switch inst.operand {
			case valA:
				e.registers[regB] = e.registers[regA] % 8
			case valB:
				e.registers[regB] = e.registers[regB] % 8
			case valC:
				e.registers[regB] = e.registers[regC] % 8
			default:
				e.registers[regB] = T(inst.operand)
			}
			e.instPtr++
		case jnz:
			if e.registers[regA] != 0 {
				e.instPtr = int(inst.operand)
			} else {
				e.instPtr++
			}
		case bxc:
			e.registers[regB] ^= e.registers[regC]
			e.instPtr++
		case out:
			var outVal T
			switch inst.operand {
			case valA:
				outVal = e.registers[regA] % 8
			case valB:
				outVal = e.registers[regB] % 8
			case valC:
				outVal = e.registers[regC] % 8
			default:
				outVal = T(inst.operand)
			}
			outputs = append(outputs, fmt.Sprintf("%d", outVal))
			e.instPtr++
		case bdv:
			switch inst.operand {
			case valA:
				e.registers[regB] = e.registers[regA] >> e.registers[regA]
			case valB:
				e.registers[regB] = e.registers[regA] >> e.registers[regB]
			case valC:
				e.registers[regB] = e.registers[regA] >> e.registers[regC]
			default:
				e.registers[regB] = e.registers[regA] >> inst.operand
			}
			e.instPtr++
		case cdv:
			switch inst.operand {
			case valA:
				e.registers[regC] = e.registers[regA] >> e.registers[regA]
			case valB:
				e.registers[regC] = e.registers[regA] >> e.registers[regB]
			case valC:
				e.registers[regC] = e.registers[regA] >> e.registers[regC]
			default:
				e.registers[regC] = e.registers[regA] >> inst.operand
			}
			e.instPtr++
		}
	}
	return strings.Join(outputs, ",")
}

func (e emulator[T]) String() string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("Register A: %d", e.registers[regA]))
	lines = append(lines, fmt.Sprintf("Register B: %d", e.registers[regB]))
	lines = append(lines, fmt.Sprintf("Register C: %d", e.registers[regC]))
	lines = append(lines, "")
	lines = append(lines, "Program:")
	lines = append(lines, "")
	for _, inst := range e.program {
		switch inst.opcode {
		case adv:
			switch inst.operand {
			case valA:
				lines = append(lines, "adv valA")
			case valB:
				lines = append(lines, "adv valB")
			case valC:
				lines = append(lines, "adv valC")
			default:
				lines = append(lines, fmt.Sprintf("adv %d", inst.operand))
			}
		case bxl:
			lines = append(lines, fmt.Sprintf("bxl %d", inst.operand))
		case bst:
			switch inst.operand {
			case valA:
				lines = append(lines, "bst valA")
			case valB:
				lines = append(lines, "bst valB")
			case valC:
				lines = append(lines, "bst valC")
			default:
				lines = append(lines, fmt.Sprintf("bst %d", inst.operand))
			}
		case jnz:
			lines = append(lines, fmt.Sprintf("jnz %d", inst.operand))
		case bxc:
			lines = append(lines, "bxc")
		case out:
			switch inst.operand {
			case valA:
				lines = append(lines, "out valA")
			case valB:
				lines = append(lines, "out valB")
			case valC:
				lines = append(lines, "out valC")
			default:
				lines = append(lines, fmt.Sprintf("out %d", inst.operand))
			}
		case bdv:
			switch inst.operand {
			case valA:
				lines = append(lines, "bdv valA")
			case valB:
				lines = append(lines, "bdv valB")
			case valC:
				lines = append(lines, "bdv valC")
			default:
				lines = append(lines, fmt.Sprintf("bdv %d", inst.operand))
			}
		case cdv:
			switch inst.operand {
			case valA:
				lines = append(lines, "cdv valA")
			case valB:
				lines = append(lines, "cdv valB")
			case valC:
				lines = append(lines, "cdv valC")
			default:
				lines = append(lines, fmt.Sprintf("cdv %d", inst.operand))
			}
		}
	}
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}

func parseEmulator[T ~int | ~int32 | ~int64](inputs []string) emulator[T] {
	matcher := regexp.MustCompile(regPattern)
	registers := [3]T{}
	for i, s := range inputs[:3] {
		match := matcher.FindStringSubmatch(s)
		if regVal, err := strconv.Atoi(match[2]); err == nil {
			registers[i] = T(regVal)
		} else {
			log.Panicf("Could not parse register value from input %q", s)
		}
	}
	matcher = regexp.MustCompile(progPattern)
	program := []instruction{}
	match := matcher.FindStringSubmatch(inputs[4])
	split := strings.Split(match[1], ",")
	for i := 0; i < len(split)-1; i += 2 {
		opcode, err := strconv.ParseUint(split[i], 10, 3)
		if err != nil {
			log.Panicf("Could not parse opcode value from input %q", inputs[4])
		}
		operand, err := strconv.ParseUint(split[i+1], 10, 3)
		if err != nil {
			log.Panicf("Could not parse operand value from input %q", inputs[4])
		}
		program = append(program, instruction{uint8(opcode), uint8(operand)})
	}
	return emulator[T]{
		registers: registers,
		program:   program,
	}
}

func ExecProgram(inputs []string) string {
	e := parseEmulator[int](inputs)
	log.Printf("Running with emulator:\n%s", e)
	return e.execute()
}
