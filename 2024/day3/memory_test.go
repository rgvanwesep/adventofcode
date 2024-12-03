package day3

import "testing"

func TestSumMul(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected uint64
	}{
		{
			[]string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			161,
		},
	}
	for _, c := range cases {
		result := SumMul(c.inputs)
		if result != c.expected {
			t.Errorf("SumMul(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestSumConditionalMul(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected uint64
	}{
		{
			[]string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			48,
		},
	}
	for _, c := range cases {
		result := SumConditionalMul(c.inputs)
		if result != c.expected {
			t.Errorf("SumConditionalMul(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
