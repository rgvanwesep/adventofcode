package day4

import "testing"

func TestCountOccurances(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			18,
		},
	}
	for _, c := range cases {
		result := CountOccurances(c.inputs)
		if result != c.expected {
			t.Errorf("CountOccurances(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestCountOccurancesX(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			9,
		},
	}
	for _, c := range cases {
		result := CountOccurancesX(c.inputs)
		if result != c.expected {
			t.Errorf("CountOccurancesX(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
