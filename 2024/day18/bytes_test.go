package day18

import "testing"

func TestCountSteps(t *testing.T) {
	cases := []struct {
		inputs       []string
		nrows, ncols int
		nInputs      int
		expected     int
	}{
		{
			inputs: []string{
				"5,4",
				"4,2",
				"4,5",
				"3,0",
				"2,1",
				"6,3",
				"2,4",
				"1,5",
				"0,6",
				"3,3",
				"2,6",
				"5,1",
				"1,2",
				"5,5",
				"2,5",
				"6,5",
				"1,4",
				"0,4",
				"6,4",
				"1,1",
				"6,1",
				"1,0",
				"0,5",
				"1,6",
				"2,0",
			},
			nrows:    7,
			ncols:    7,
			nInputs:  12,
			expected: 22,
		},
	}
	for _, c := range cases {
		result := CountSteps(c.inputs, c.nrows, c.ncols, c.nInputs)
		if result != c.expected {
			t.Errorf("CountSteps(%q, %d, %d, %d) == %d, expected %d",
				c.inputs,
				c.nrows,
				c.ncols,
				c.nInputs,
				result,
				c.expected,
			)
		}
	}
}
