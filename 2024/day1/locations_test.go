package day1

import "testing"

func TestSumDistances(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected uint64
	}{
		{
			[]string{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			11,
		},
	}
	for _, c := range cases {
		result := SumDistances(c.inputs)
		if result != c.expected {
			t.Errorf("Sum(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestCalcSimilarity(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected uint64
	}{
		{
			[]string{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			31,
		},
	}
	for _, c := range cases {
		result := CalcSimilarity(c.inputs)
		if result != c.expected {
			t.Errorf("Sum(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}