package day19

import "testing"

func TestCountSteps(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"r, wr, b, g, bwu, rb, gb, br",
				"",
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			expected: 6,
		},
	}
	for _, c := range cases {
		result := CountPossible(c.inputs)
		if result != c.expected {
			t.Errorf("CountPossible(%q) == %d, expected %d",
				c.inputs,
				result,
				c.expected,
			)
		}
	}
}
