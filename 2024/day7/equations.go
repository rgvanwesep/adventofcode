package day7

import (
	"log"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
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
	if len(ops) != len(e.terms)-1 {
		log.Panicf("Mismatch between ops and terms lengths: %d != %d", len(ops), len(e.terms)-1)
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

func CalcDigits(x int) int {
	n := 0
	for x > 0 {
		n++
		x /= 10
	}
	return n
}

type ConcatOp struct{}

func (ConcatOp) Eval(x, y int) int {
	nDigitsY := CalcDigits(y)
	factor := 1
	for i := 0; i < nDigitsY; i++ {
		factor *= 10
	}
	return x*factor + y
}

var _ BinaryOp = ConcatOp{}

func GenerateOps(n int) [][]BinaryOp {
	if n == 1 {
		return [][]BinaryOp{{AddOp{}}, {MultiplyOp{}}}
	}
	ops := make([][]BinaryOp, 0)
	prevOps := GenerateOps(n - 1)
	for _, o := range prevOps {
		add := append(slices.Clone(o), AddOp{})
		multiply := append(o, MultiplyOp{})
		ops = append(ops, add, multiply)
	}
	return ops
}

func GenerateOpsWithConcat(n int) [][]BinaryOp {
	if n == 1 {
		return [][]BinaryOp{{AddOp{}}, {MultiplyOp{}}, {ConcatOp{}}}
	}
	ops := make([][]BinaryOp, 0)
	prevOps := GenerateOpsWithConcat(n - 1)
	for _, o := range prevOps {
		add := append(slices.Clone(o), AddOp{})
		multiply := append(slices.Clone(o), MultiplyOp{})
		concat := append(slices.Clone(o), ConcatOp{})
		ops = append(ops, add, multiply, concat)
	}
	return ops
}

func ParseEquations(inputs []string) []Equation {
	equations := make([]Equation, 0)
	for _, input := range inputs {
		split := strings.Split(input, ": ")
		if len(split) != 2 {
			log.Panicf("Could not split input %q", input)
		}
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

func SumCorrectedWithConcat(inputs []string) int {
	sum := 0
	maxSum := big.NewInt(0)
	equations := ParseEquations(inputs)
	for _, equation := range equations {
		maxSum.Add(maxSum, big.NewInt(int64(equation.result)))
		opsCombos := GenerateOpsWithConcat(len(equation.terms) - 1)
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

func GetPossibleResults(termStr string) []int {
	split := strings.Split(termStr, " ")
	terms := make([]int, 0)
	for _, s := range split {
		term, err := strconv.Atoi(s)
		if err != nil {
			log.Panicf("Could not parse term %q from input %q", s, termStr)
		}
		terms = append(terms, term)
	}

	possibleResults := []int{terms[0] + terms[1], terms[0] * terms[1]}
	for i := 2; i < len(terms); i++ {
		newPossibleResults := make([]int, 0)
		for _, r := range possibleResults {
			newPossibleResults = append(newPossibleResults, r+terms[i], r*terms[i])
		}
		possibleResults = newPossibleResults
	}
	return possibleResults
}

func SumCorrectedSimple(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		split := strings.Split(input, ": ")
		if len(split) != 2 {
			log.Panicf("Could not split input %q", input)
		}
		resultStr, termStr := split[0], split[1]
		result, err := strconv.Atoi(resultStr)
		if err != nil {
			log.Panicf("Could not parse result from input %q", input)
		}

		possibleResults := GetPossibleResults(termStr)
		for _, r := range possibleResults {
			if r == result {
				sum += r
				break
			}
		}
	}
	return sum
}
