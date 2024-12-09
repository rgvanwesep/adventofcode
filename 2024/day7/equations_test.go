package day7

import (
	"math/big"
	"testing"
)

func TestSumCorrected(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected *big.Int
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
			big.NewInt(3749),
		},
	}
	for _, c := range cases {
		result := SumCorrected(c.inputs)
		if result.Cmp(c.expected) != 0 {
			t.Errorf("SumCorrected(%v) == %v, expected %v", c.inputs, result, c.expected)
		}
	}
}
