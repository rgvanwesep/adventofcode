package day7

import (
	"log"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	terms  []int
}

func (e Equation) EvalCheckWith(ops []BinaryOp) bool {
	result := ops[0].Eval(e.terms[0], e.terms[1])
	for i := 1; i < len(e.terms)-1; i++ {
		result = ops[i].Eval(result, e.terms[i+1])
	}
	return result == e.result
}

type BinaryOp interface {
	Eval(x, y int) int
}

type AddOp struct{}

func (AddOp) Eval(x, y int) int {
	return x + y
}

var _ BinaryOp = AddOp{}

type MultiplyOp struct{}

func (MultiplyOp) Eval(x, y int) int {
	return x * y
}

var _ BinaryOp = MultiplyOp{}

func GenerateOps(n int) [][]BinaryOp {
	if n == 1 {
		return [][]BinaryOp{{AddOp{}}, {MultiplyOp{}}}
	}
	ops := make([][]BinaryOp, 0)
	prevOps := GenerateOps(n - 1)
	for _, o := range prevOps {
		add := append(o, AddOp{})
		multiply := append(o, MultiplyOp{})
		ops = append(ops, add, multiply)
	}
	return ops
}

func ParseEquations(inputs []string) []Equation {
	equations := make([]Equation, 0)
	for _, input := range inputs {
		split := strings.Split(input, ": ")
		resultStr, termStr := split[0], split[1]
		result, err := strconv.Atoi(resultStr)
		if err != nil {
			log.Panicf("Could not parse result from input %q", input)
		}
		split = strings.Split(termStr, " ")
		terms := make([]int, 0)
		for _, s := range split {
			term, err := strconv.Atoi(s)
			if err != nil {
				log.Panicf("Could not parse term %q from input %q", s, input)
			}
			terms = append(terms, term)
		}
		equations = append(equations, Equation{result, terms})
	}
	return equations
}

func SumCorrected(inputs []string) int {
	sum := 0
	equations := ParseEquations(inputs)
	for _, equation := range equations {
		opsCombos := GenerateOps(len(equation.terms) - 1)
		for _, ops := range opsCombos {
			if ok := equation.EvalCheckWith(ops); ok {
				sum += equation.result
				break
			}
		}
	}
	return sum
}
