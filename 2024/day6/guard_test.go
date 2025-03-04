package day6

import "testing"

func TestCountVisited(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			41,
		},
	}
	for _, c := range cases {
		result := CountVisited(c.inputs)
		if result != c.expected {
			t.Errorf("CountVisited(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestCountCyclingObstructions(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			6,
		},
	}
	for _, c := range cases {
		result := CountCyclingObstructions(c.inputs)
		if result != c.expected {
			t.Errorf("CountCyclingObstructions(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
