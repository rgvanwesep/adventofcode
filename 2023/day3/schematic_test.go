package day3

import "testing"

func TestExtend(t *testing.T) {
	cases := []struct {
		cRange   CoordinateRange
		coord    Coordinate
		expected CoordinateRange
	}{
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}},
			Coordinate{0, 0},
			CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}},
			Coordinate{1, 0},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 0}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}},
			Coordinate{0, 1},
			CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}},
			Coordinate{1, 1},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{0, 0},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{1, 1},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{0, 1},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{1, 0},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{2, 0},
			CoordinateRange{Coordinate{0, 0}, Coordinate{2, 1}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{0, 2},
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 2}},
		},
		{
			CoordinateRange{Coordinate{0, 0}, Coordinate{1, 1}},
			Coordinate{2, 2},
			CoordinateRange{Coordinate{0, 0}, Coordinate{2, 2}},
		},
		{
			CoordinateRange{Coordinate{1, 1}, Coordinate{1, 2}},
			Coordinate{1, 0},
			CoordinateRange{Coordinate{1, 0}, Coordinate{1, 2}},
		},
		{
			CoordinateRange{Coordinate{1, 1}, Coordinate{1, 2}},
			Coordinate{0, 1},
			CoordinateRange{Coordinate{0, 1}, Coordinate{1, 2}},
		},
		{
			CoordinateRange{Coordinate{1, 1}, Coordinate{1, 2}},
			Coordinate{1, 3},
			CoordinateRange{Coordinate{1, 1}, Coordinate{1, 3}},
		},
	}
	for _, c := range cases {
		c.cRange.Extend(c.coord)
		if c.cRange != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, c.cRange)
		}
	}
}

func TestAppendDigit(t *testing.T) {
	cases := []struct {
		partNumber PartNumber
		digit      int
		coordinate Coordinate
		expected   PartNumber
	}{
		{
			PartNumber{1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
			1,
			Coordinate{0, 1},
			PartNumber{11, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
		},
	}
	for _, c := range cases {
		c.partNumber.AppendDigit(c.digit, c.coordinate)
		if c.partNumber != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, c.partNumber)
		}
	}
}

func TestEquals(t *testing.T) {
	cases := []struct {
		schematic1 Schematic
		schematic2 Schematic
		expected   bool
	}{
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			true,
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			true,
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {2, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			false,
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			false,
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
					{0, 1}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			false,
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{
					{0, 0}: {'a', Coordinate{0, 0}},
				},
			},
			false,
		},
	}
	for _, c := range cases {
		if c.schematic1.Equals(&c.schematic2) != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, c.schematic1.Equals(&c.schematic2))
		}
	}
}

func TestAddDigit(t *testing.T) {
	cases := []struct {
		schematic  Schematic
		digit      int
		coordinate Coordinate
		expected   Schematic
	}{
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			1,
			Coordinate{0, 0},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			1,
			Coordinate{0, 1},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {11, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
					{0, 1}: {11, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {11, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
					{0, 1}: {11, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			2,
			Coordinate{0, 2},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {112, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 2}}},
					{0, 1}: {112, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 2}}},
					{0, 2}: {112, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 2}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
		},
	}
	for _, c := range cases {
		c.schematic.AddDigit(c.digit, c.coordinate)
		if !c.schematic.Equals(&c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, c.schematic)
		}
	}
}

func TestDiff(t *testing.T) {
	cases := []struct {
		schematic1 Schematic
		schematic2 Schematic
		expected   map[Coordinate]bool
	}{
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			map[Coordinate]bool{},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			map[Coordinate]bool{
				{0, 0}: true,
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols: map[Coordinate]Symbol{
					{0, 0}: {'a', Coordinate{0, 0}},
				},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols:     map[Coordinate]Symbol{},
			},
			map[Coordinate]bool{
				{0, 0}: true,
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols: map[Coordinate]Symbol{
					{0, 0}: {'a', Coordinate{0, 0}},
				},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols: map[Coordinate]Symbol{
					{0, 0}: {'b', Coordinate{0, 0}},
				},
			},
			map[Coordinate]bool{
				{0, 0}: true,
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols: map[Coordinate]Symbol{
					{0, 0}: {'a', Coordinate{0, 0}},
				},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{},
				Symbols: map[Coordinate]Symbol{
					{0, 1}: {'a', Coordinate{0, 1}},
				},
			},
			map[Coordinate]bool{
				{0, 0}: true,
				{0, 1}: true,
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {2, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			map[Coordinate]bool{
				{0, 0}: true,
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 1}: {1, CoordinateRange{Coordinate{0, 1}, Coordinate{0, 1}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
			map[Coordinate]bool{
				{0, 0}: true,
				{0, 1}: true,
			},
		},
	}
	for _, c := range cases {
		diff := c.schematic1.Diff(&c.schematic2)
		if len(diff) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, diff)
		}
		for i, coord := range diff {
			if coord != c.expected[i] {
				t.Errorf("Expected %v, got %v", c.expected, diff)
			}
		}
	}
}

func TestNewSchematic(t *testing.T) {
	cases := []struct {
		inputs    []string
		schematic Schematic
	}{
		{
			[]string{"1"},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
		},
		{
			[]string{"1", "2"},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
					{1, 0}: {2, CoordinateRange{Coordinate{1, 0}, Coordinate{1, 0}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
		},
		{
			[]string{"12", "34"},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {12, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
					{0, 1}: {12, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 1}}},
					{1, 0}: {34, CoordinateRange{Coordinate{1, 0}, Coordinate{1, 1}}},
					{1, 1}: {34, CoordinateRange{Coordinate{1, 0}, Coordinate{1, 1}}},
				},
				Symbols: map[Coordinate]Symbol{},
			},
		},
		{
			[]string{"...#....", "..123...", "........", "..765..."},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{1, 2}: {123, CoordinateRange{Coordinate{1, 2}, Coordinate{1, 4}}},
					{1, 3}: {123, CoordinateRange{Coordinate{1, 2}, Coordinate{1, 4}}},
					{1, 4}: {123, CoordinateRange{Coordinate{1, 2}, Coordinate{1, 4}}},
					{3, 2}: {765, CoordinateRange{Coordinate{3, 2}, Coordinate{3, 4}}},
					{3, 3}: {765, CoordinateRange{Coordinate{3, 2}, Coordinate{3, 4}}},
					{3, 4}: {765, CoordinateRange{Coordinate{3, 2}, Coordinate{3, 4}}},
				},
				Symbols: map[Coordinate]Symbol{
					{0, 3}: {'#', Coordinate{0, 3}},
				},
			},
		},
		{
			[]string{
				"467..114..\n",
				"...*......\n",
				"..35..633.\n",
				"......#...\n",
				"617*......\n",
				".....+.58.\n",
				"..592.....\n",
				"......755.\n",
				"...$.*....\n",
				".664.598..",
			},
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {467, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 2}}},
					{0, 1}: {467, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 2}}},
					{0, 2}: {467, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 2}}},
					{0, 5}: {114, CoordinateRange{Coordinate{0, 5}, Coordinate{0, 7}}},
					{0, 6}: {114, CoordinateRange{Coordinate{0, 5}, Coordinate{0, 7}}},
					{0, 7}: {114, CoordinateRange{Coordinate{0, 5}, Coordinate{0, 7}}},
					{2, 2}: {35, CoordinateRange{Coordinate{2, 2}, Coordinate{2, 3}}},
					{2, 3}: {35, CoordinateRange{Coordinate{2, 2}, Coordinate{2, 3}}},
					{2, 6}: {633, CoordinateRange{Coordinate{2, 6}, Coordinate{2, 8}}},
					{2, 7}: {633, CoordinateRange{Coordinate{2, 6}, Coordinate{2, 8}}},
					{2, 8}: {633, CoordinateRange{Coordinate{2, 6}, Coordinate{2, 8}}},
					{4, 0}: {617, CoordinateRange{Coordinate{4, 0}, Coordinate{4, 2}}},
					{4, 1}: {617, CoordinateRange{Coordinate{4, 0}, Coordinate{4, 2}}},
					{4, 2}: {617, CoordinateRange{Coordinate{4, 0}, Coordinate{4, 2}}},
					{5, 7}: {58, CoordinateRange{Coordinate{5, 7}, Coordinate{5, 8}}},
					{5, 8}: {58, CoordinateRange{Coordinate{5, 7}, Coordinate{5, 8}}},
					{6, 2}: {592, CoordinateRange{Coordinate{6, 2}, Coordinate{6, 4}}},
					{6, 3}: {592, CoordinateRange{Coordinate{6, 2}, Coordinate{6, 4}}},
					{6, 4}: {592, CoordinateRange{Coordinate{6, 2}, Coordinate{6, 4}}},
					{7, 6}: {755, CoordinateRange{Coordinate{7, 6}, Coordinate{7, 8}}},
					{7, 7}: {755, CoordinateRange{Coordinate{7, 6}, Coordinate{7, 8}}},
					{7, 8}: {755, CoordinateRange{Coordinate{7, 6}, Coordinate{7, 8}}},
					{9, 1}: {664, CoordinateRange{Coordinate{9, 1}, Coordinate{9, 3}}},
					{9, 2}: {664, CoordinateRange{Coordinate{9, 1}, Coordinate{9, 3}}},
					{9, 3}: {664, CoordinateRange{Coordinate{9, 1}, Coordinate{9, 3}}},
					{9, 5}: {598, CoordinateRange{Coordinate{9, 5}, Coordinate{9, 7}}},
					{9, 6}: {598, CoordinateRange{Coordinate{9, 5}, Coordinate{9, 7}}},
					{9, 7}: {598, CoordinateRange{Coordinate{9, 5}, Coordinate{9, 7}}},
				},
				Symbols: map[Coordinate]Symbol{
					{1, 3}: {'*', Coordinate{1, 3}},
					{3, 6}: {'#', Coordinate{3, 6}},
					{4, 3}: {'*', Coordinate{4, 3}},
					{5, 5}: {'+', Coordinate{5, 5}},
					{8, 3}: {'$', Coordinate{8, 3}},
					{8, 5}: {'*', Coordinate{8, 5}},
				},
			},
		},
	}
	for _, c := range cases {
		schematic := NewSchematic(c.inputs)
		if !schematic.Equals(&c.schematic) {
			t.Errorf("Expected %v, got %v\nDiff: %v", &c.schematic, schematic, schematic.Diff(&c.schematic))
		}
	}
}

func TestNeighbors(t *testing.T) {
	cases := []struct {
		coordinate Coordinate
		expected   map[Coordinate]bool
	}{
		{
			Coordinate{1, 1},
			map[Coordinate]bool{
				{0, 0}: true,
				{0, 1}: true,
				{0, 2}: true,
				{1, 0}: true,
				{1, 2}: true,
				{2, 0}: true,
				{2, 1}: true,
				{2, 2}: true,
			},
		},
	}
	for _, c := range cases {
		neighbors := c.coordinate.Neighbors()
		if len(neighbors) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, neighbors)
		}
		for neighbor := range neighbors {
			if !c.expected[neighbor] {
				t.Errorf("Expected %v, got %v", c.expected, neighbors)
			}
		}
	}
}

func TestAdjacentPartNumbers(t *testing.T) {
	cases := []struct {
		schematic Schematic
		coord     Coordinate
		expected  map[PartNumber]bool
	}{
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
				},
				Symbols: map[Coordinate]Symbol{
					{1, 1}: {'a', Coordinate{1, 1}},
				},
			},
			Coordinate{1, 1},
			map[PartNumber]bool{
				{1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}}: true,
			},
		},
		{
			Schematic{
				PartNumbers: map[Coordinate]PartNumber{
					{0, 0}: {1, CoordinateRange{Coordinate{0, 0}, Coordinate{0, 0}}},
					{1, 0}: {2, CoordinateRange{Coordinate{1, 0}, Coordinate{1, 0}}},
				},
				Symbols: map[Coordinate]Symbol{
					{2, 1}: {'a', Coordinate{2, 1}},
				},
			},
			Coordinate{2, 1},
			map[PartNumber]bool{
				{2, CoordinateRange{Coordinate{1, 0}, Coordinate{1, 0}}}: true,
			},
		},
	}
	for _, c := range cases {
		partNumbers := c.schematic.AdjacentPartNumbers(c.coord)
		if len(partNumbers) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, partNumbers)
		}
		for partNumber := range partNumbers {
			if !c.expected[partNumber] {
				t.Errorf("Expected %v, got %v", c.expected, partNumbers)
			}
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"467..114..\n",
				"...*......\n",
				"..35..633.\n",
				"......#...\n",
				"617*......\n",
				".....+.58.\n",
				"..592.....\n",
				"......755.\n",
				"...$.*....\n",
				".664.598..",
			},
			4361,
		},
	}
	for _, c := range cases {
		sum := Sum(c.inputs)
		if sum != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, sum)
		}
	}
}

func TestSumGearRatios(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"467..114..\n",
				"...*......\n",
				"..35..633.\n",
				"......#...\n",
				"617*......\n",
				".....+.58.\n",
				"..592.....\n",
				"......755.\n",
				"...$.*....\n",
				".664.598..",
			},
			467835,
		},
	}
	for _, c := range cases {
		sum := SumGearRatios(c.inputs)
		if sum != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, sum)
		}
	}
}
