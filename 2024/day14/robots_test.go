package day14

import "testing"

func TestCalcSafetyFactor(t *testing.T) {
	cases := []struct {
		inputs       []string
		nrows, ncols int
		nIter        int
		expected     int
	}{
		{
			inputs: []string{
				"p=0,4 v=3,-3",
				"p=6,3 v=-1,-3",
				"p=10,3 v=-1,2",
				"p=2,0 v=2,-1",
				"p=0,0 v=1,3",
				"p=3,0 v=-2,-2",
				"p=7,6 v=-1,-3",
				"p=3,0 v=-1,-2",
				"p=9,3 v=2,3",
				"p=7,3 v=-1,2",
				"p=2,4 v=2,-3",
				"p=9,5 v=-3,-3",
			},
			nrows:    7,
			ncols:    11,
			nIter:    100,
			expected: 12,
		},
	}
	for _, c := range cases {
		result := CalcSafetyFactor(c.inputs, c.nrows, c.ncols, c.nIter)
		if result != c.expected {
			t.Errorf(
				"CalcSafetyFactor(%q, %d, %d, %d) == %d, expected %d",
				c.inputs,
				c.nrows,
				c.ncols,
				c.nIter,
				result,
				c.expected,
			)
		}
	}
}
