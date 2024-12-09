package day7

import (
	"log"
	"math/big"
	"strconv"
	"strings"
)

const (
	MaxUint = ^uint(0)
	MaxInt = int(MaxUint >> 1)
)

type Equation struct {
	result int
	terms  []int
}

func NewEquation(result int, terms []int) Equation {
	if len(terms) < 2 {
		log.Panic("Equation must have two or more terms")
	}
	return Equation{result, terms}
}

func (e Equation) EvalCheckWith(ops []BinaryOp) bool {
	if len(ops) != len(e.terms) - 1 {
		log.Panicf("Mismatch between ops and terms lengths: %d != %d", len(ops), len(e.terms) - 1)
	}
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
		equations = append(equations, NewEquation(result, terms))
	}
	return equations
}

func SumCorrected(inputs []string) int {
	sum := 0
	maxSum := big.NewInt(0)
	equations := ParseEquations(inputs)
	for _, equation := range equations {
		maxSum.Add(maxSum, big.NewInt(int64(equation.result)))
		opsCombos := GenerateOps(len(equation.terms) - 1)
		for _, ops := range opsCombos {
			if ok := equation.EvalCheckWith(ops); ok {
				sum += equation.result
				break
			}
		}
	}
	if maxSum.Cmp(big.NewInt(int64(MaxInt))) <= 0 {
		log.Print("Int is large enough")
	} else {
		log.Print("Int may not be large enough")
	}
	return sum
}
