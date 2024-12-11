package day11

import "testing"

func TestCountPebbles(t *testing.T) {
	cases := []struct {
		inputs   []string
		nBlinks  int
		expected int
	}{
		{[]string{"125 17"}, 25, 55312},
	}
	for _, c := range cases {
		result := CountPebbles(c.inputs, c.nBlinks)
		if result != c.expected {
			t.Errorf("CountPebbles(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
