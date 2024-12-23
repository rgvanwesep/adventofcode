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

func TestIsValidWithInput(t *testing.T) {
	cases := []struct {
		input    byte
		from, to string
		expected bool
	}{
		{
			input:    '<',
			from:     "AAA",
			to:       "^AA",
			expected: true,
		},
		{
			input:    'v',
			from:     "^AA",
			to:       "vAA",
			expected: true,
		},
		{
			input:    'A',
			from:     "vAA",
			to:       "v>A",
			expected: true,
		},
		{
			input:    '<',
			from:     "v>A",
			to:       "<>A",
			expected: true,
		},
		{
			input:    'A',
			from:     "<>A",
			to:       "<vA",
			expected: true,
		},
		{
			input:    'A',
			from:     "<vA",
			to:       "<<A",
			expected: true,
		},
		{
			input:    '>',
			from:     "<<A",
			to:       "v<A",
			expected: true,
		},
	}
	for _, c := range cases {
		result := isValidWithInput(c.input, c.from, c.to)
		if result != c.expected {
			t.Errorf("isValidWithInput(%q, %q, %q) == %t, expected %t",
				c.input, c.from, c.to, result, c.expected,
			)
		}
	}
}
