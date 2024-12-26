package day25

import "testing"

func TestCountFits(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"#####",
				".####",
				".####",
				".####",
				".#.#.",
				".#...",
				".....",
				"",
				"#####",
				"##.##",
				".#.##",
				"...##",
				"...#.",
				"...#.",
				".....",
				"",
				".....",
				"#....",
				"#....",
				"#...#",
				"#.#.#",
				"#.###",
				"#####",
				"",
				".....",
				".....",
				"#.#..",
				"###..",
				"###.#",
				"###.#",
				"#####",
				"",
				".....",
				".....",
				".....",
				"#....",
				"#.#..",
				"#.#.#",
				"#####",
			},
			expected: 3,
		},
	}
	for _, c := range cases {
		result := CountFits(c.inputs)
		if result != c.expected {
			t.Errorf("CountFits(%q) == %d, expected %d",
				c.inputs, result, c.expected,
			)
		}
	}
}
