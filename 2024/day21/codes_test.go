package day21

import "testing"

func TestCalcComplexity(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"029A",
				"980A",
				"179A",
				"456A",
				"379A",
			},
			expected: 126384,
		},
		{
			inputs: []string{
				"789A",
				"968A",
				"286A",
				"349A",
				"170A",
			},
			expected: 176650,
		},
	}
	for _, c := range cases {
		result := CalcComplexity(c.inputs)
		if result != c.expected {
			t.Errorf("CalcComplexity(%q) == %d, expected %d",
				c.inputs,
				result,
				c.expected,
			)
		}
	}
}

func TestGetShortestSequenceLength(t *testing.T) {
	cases := []struct {
		input               string
		nDirectionalKeypads int
		expected            int
	}{
		{
			input:               "029A",
			nDirectionalKeypads: 3,
			expected:            68,
		},
		{
			input:               "980A",
			nDirectionalKeypads: 3,
			expected:            60,
		},
		{
			input:               "179A",
			nDirectionalKeypads: 3,
			expected:            68,
		},
		{
			input:               "456A",
			nDirectionalKeypads: 3,
			expected:            64,
		},
		{
			input:               "379A",
			nDirectionalKeypads: 3,
			expected:            64,
		},
	}
	for _, c := range cases {
		result := getShortestSequenceLength(c.input, c.nDirectionalKeypads)
		if result != c.expected {
			t.Errorf("getShortestSequenceLength(%q, %d) == %d, expected %d",
				c.input,
				c.nDirectionalKeypads,
				result,
				c.expected,
			)
		}
	}
}
