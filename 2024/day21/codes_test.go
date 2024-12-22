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
