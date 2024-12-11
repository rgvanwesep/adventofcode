package day11

import "testing"

func TestCountPebbles(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{[]string{"125 17"}, 55312},
	}
	for _, c := range cases {
		result := CountPebbles(c.inputs)
		if result != c.expected {
			t.Errorf("CountPebbles(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
