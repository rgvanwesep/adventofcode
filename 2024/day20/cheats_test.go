package day20

import "testing"

func TestCountCheatsBySavings(t *testing.T) {
	cases := []struct {
		inputs    []string
		maxCost   int
		threshold int
		expected  map[int]int
	}{
		{
			inputs: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			maxCost:   2,
			threshold: 1,
			expected: map[int]int{
				2:  14,
				4:  14,
				6:  2,
				8:  4,
				10: 2,
				12: 3,
				20: 1,
				36: 1,
				38: 1,
				40: 1,
				64: 1,
			},
		},
		{
			inputs: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			maxCost:   20,
			threshold: 50,
			expected: map[int]int{
				50: 32,
				52: 31,
				54: 29,
				56: 39,
				58: 25,
				60: 23,
				62: 20,
				64: 19,
				66: 12,
				68: 14,
				70: 12,
				72: 22,
				74: 4,
				76: 3,
			},
		},
	}
	for _, c := range cases {
		result := countCheatsBySavings(c.inputs, c.maxCost, c.threshold)
		if len(result) != len(c.expected) {
			t.Errorf("len(countCheatsBySavings(\n%s, %d, %d)) == %d, expected %d",
				parseGrid(c.inputs), c.maxCost, c.threshold, len(result), len(c.expected),
			)
		} else {
			for k, v := range result {
				if v != c.expected[k] {
					t.Errorf("countCheatsBySavings(\n%s, %d, %d)[%d] == %d, expected %d",
						parseGrid(c.inputs), c.maxCost, c.threshold, k, v, c.expected[k],
					)
				}
			}
		}
	}
}
