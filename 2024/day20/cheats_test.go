package day20

import "testing"

func TestCountCheatsBySavings(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected map[int]int
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
	}
	for _, c := range cases {
		result := countCheatsBySavings(c.inputs)
		if len(result) != len(c.expected) {
			t.Errorf("len(countCheatsBySavings(\n%s)) == %d, expected %d", parseGrid(c.inputs), len(result), len(c.expected))
		} else {
			for k, v := range result {
				if v != c.expected[k] {
					t.Errorf("countCheatsBySavings(\n%s)[%d] == %d, expected %d", parseGrid(c.inputs), k, v, c.expected[k])
				}
			}
		}
	}
}
