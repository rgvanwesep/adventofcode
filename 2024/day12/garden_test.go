package day12

import "testing"

func TestSumFencePrice(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"AAAA",
				"BBCD",
				"BBCC",
				"EEEC",
			},
			140,
		},
		{
			[]string{
				"OOOOO",
				"OXOXO",
				"OOOOO",
				"OXOXO",
				"OOOOO",
			},
			772,
		},
		{
			[]string{
				"RRRRIICCFF",
				"RRRRIICCCF",
				"VVRRRCCFFF",
				"VVRCCCJFFF",
				"VVVVCJJCFE",
				"VVIVCCJJEE",
				"VVIIICJJEE",
				"MIIIIIJJEE",
				"MIIISIJEEE",
				"MMMISSJEEE",
			},
			1930,
		},
	}
	for _, c := range cases {
		result := SumFencePrice(c.inputs)
		if result != c.expected {
			t.Errorf("SumFencePrice(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestSumFencePriceDiscount(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"AAAA",
				"BBCD",
				"BBCC",
				"EEEC",
			},
			80,
		},
		{
			[]string{
				"OOOOO",
				"OXOXO",
				"OOOOO",
				"OXOXO",
				"OOOOO",
			},
			436,
		},
		{
			[]string{
				"EEEEE",
				"EXXXX",
				"EEEEE",
				"EXXXX",
				"EEEEE",
			},
			236,
		},
		{
			[]string{
				"AAAAAA",
				"AAABBA",
				"AAABBA",
				"ABBAAA",
				"ABBAAA",
				"AAAAAA",
			},
			368,
		},
		{
			[]string{
				"RRRRIICCFF",
				"RRRRIICCCF",
				"VVRRRCCFFF",
				"VVRCCCJFFF",
				"VVVVCJJCFE",
				"VVIVCCJJEE",
				"VVIIICJJEE",
				"MIIIIIJJEE",
				"MIIISIJEEE",
				"MMMISSJEEE",
			},
			1206,
		},
	}
	for _, c := range cases {
		result := SumFencePriceDiscount(c.inputs)
		if result != c.expected {
			t.Errorf("SumFencePriceDiscount(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
