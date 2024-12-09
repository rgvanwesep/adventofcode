package day7

import "testing"

func TestSumCorrected(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			3749,
		},
		{
			[]string{
				"6: 1 2 3",
			},
			6,
		},
		{
			[]string{
				"6: 1 3 2",
			},
			6,
		},
		{
			[]string{
				"9: 1 2 3",
			},
			9,
		},
		{
			[]string{
				"9: 1 3 2",
			},
			0,
		},
		{
			[]string{
				"5: 1 2 3",
			},
			5,
		},
		{
			[]string{
				"5: 1 3 2",
			},
			5,
		},
		{
			[]string{
				"1: 1 2 3",
				"2: 1 2 3",
				"3: 1 2 3",
				"4: 1 2 3",
				"7: 1 2 3",
				"8: 1 2 3",
				"10: 1 2 3",
			},
			0,
		},
	}
	for _, c := range cases {
		result := SumCorrected(c.inputs)
		if result != c.expected {
			t.Errorf("SumCorrected(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestEvalCheckWith(t *testing.T) {
	cases := []struct {
		e        Equation
		ops      []BinaryOp
		expected bool
	}{
		{NewEquation(3, []int{1, 2}), []BinaryOp{AddOp{}}, true},
		{NewEquation(2, []int{1, 2}), []BinaryOp{AddOp{}}, false},
		{NewEquation(4, []int{1, 2}), []BinaryOp{AddOp{}}, false},
		{NewEquation(3, []int{1, 2}), []BinaryOp{MultiplyOp{}}, false},
		{NewEquation(2, []int{1, 2}), []BinaryOp{MultiplyOp{}}, true},
		{NewEquation(1, []int{1, 2}), []BinaryOp{MultiplyOp{}}, false},
		{NewEquation(6, []int{1, 2, 3}), []BinaryOp{AddOp{}, AddOp{}}, true},
		{NewEquation(5, []int{1, 2, 3}), []BinaryOp{AddOp{}, AddOp{}}, false},
		{NewEquation(7, []int{1, 2, 3}), []BinaryOp{AddOp{}, AddOp{}}, false},
		{NewEquation(9, []int{1, 2, 3}), []BinaryOp{AddOp{}, MultiplyOp{}}, true},
		{NewEquation(8, []int{1, 2, 3}), []BinaryOp{AddOp{}, MultiplyOp{}}, false},
		{NewEquation(10, []int{1, 2, 3}), []BinaryOp{AddOp{}, MultiplyOp{}}, false},
		{NewEquation(5, []int{1, 2, 3}), []BinaryOp{MultiplyOp{}, AddOp{}}, true},
		{NewEquation(4, []int{1, 2, 3}), []BinaryOp{MultiplyOp{}, AddOp{}}, false},
		{NewEquation(6, []int{1, 2, 3}), []BinaryOp{MultiplyOp{}, AddOp{}}, false},
		{NewEquation(6, []int{1, 2, 3}), []BinaryOp{MultiplyOp{}, MultiplyOp{}}, true},
		{NewEquation(5, []int{1, 2, 3}), []BinaryOp{MultiplyOp{}, MultiplyOp{}}, false},
		{NewEquation(7, []int{1, 2, 3}), []BinaryOp{MultiplyOp{}, MultiplyOp{}}, false},
	}
	for _, c := range cases {
		result := c.e.EvalCheckWith(c.ops)
		if result != c.expected {
			t.Errorf("%v.EvalCheckWith(%v) == %v, expected %v", c.e, c.ops, result, c.expected)
		}
	}
}

func TestGenerateOps(t *testing.T) {
	cases := []struct {
		n        int
		expected [][]BinaryOp
	}{
		{
			1,
			[][]BinaryOp{{AddOp{}}, {MultiplyOp{}}},
		},
		{
			2,
			[][]BinaryOp{
				{AddOp{}, AddOp{}},
				{AddOp{}, MultiplyOp{}},
				{MultiplyOp{}, AddOp{}},
				{MultiplyOp{}, MultiplyOp{}},
			},
		},
		{
			3,
			[][]BinaryOp{
				{AddOp{}, AddOp{}, AddOp{}},
				{AddOp{}, AddOp{}, MultiplyOp{}},
				{AddOp{}, MultiplyOp{}, AddOp{}},
				{AddOp{}, MultiplyOp{}, MultiplyOp{}},
				{MultiplyOp{}, AddOp{}, AddOp{}},
				{MultiplyOp{}, AddOp{}, MultiplyOp{}},
				{MultiplyOp{}, MultiplyOp{}, AddOp{}},
				{MultiplyOp{}, MultiplyOp{}, MultiplyOp{}},
			},
		},
	}
	for _, c := range cases {
		result := GenerateOps(c.n)
		for i, ops := range c.expected {
			for j, o := range ops {
				if result[i][j] != o {
					t.Errorf("GenerateOps(%d) == %v, expected %v", c.n, result, c.expected)
				}
			}
		}
	}
}
