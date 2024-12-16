package day15

import "testing"

func TestSumCoordinates(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"########",
				"#..O.O.#",
				"##@.O..#",
				"#...O..#",
				"#.#.O..#",
				"#...O..#",
				"#......#",
				"########",
				"",
				"<^^>>>vv<v>>v<<",
			},
			expected: 2028,
		},
		{
			inputs: []string{
				"##########",
				"#..O..O.O#",
				"#......O.#",
				"#.OO..O.O#",
				"#..O@..O.#",
				"#O#..O...#",
				"#O..O..O.#",
				"#.OO.O.OO#",
				"#....O...#",
				"##########",
				"",
				"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
				"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
				"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
				"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
				"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
				"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
				">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
				"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
				"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
				"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
			},
			expected: 10092,
		},
	}
	for _, c := range cases {
		result := SumCoordinates(c.inputs)
		if result != c.expected {
			t.Errorf("SumCoordinates(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
