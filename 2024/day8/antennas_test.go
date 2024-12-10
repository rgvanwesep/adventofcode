package day8

import "testing"

func TestCountAntiNodes(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"..........",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			2,
		},
		{
			[]string{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"........a.",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			4,
		},
		{
			[]string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			},
			14,
		},
	}
	for _, c := range cases {
		result := CountAntiNodes(c.inputs)
		if result != c.expected {
			t.Errorf("CountAntiNodes(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
