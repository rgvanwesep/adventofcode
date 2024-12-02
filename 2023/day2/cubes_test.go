package day2

import "testing"

func TestNewSet(t *testing.T) {
	cases := []struct {
		input    string
		expected Set
	}{
		{"4 red, 3 blue", Set{4, 0, 3}},
		{"10 green, 4 red, 6 blue", Set{4, 10, 6}},
		{"11 red", Set{11, 0, 0}},
	}
	for _, c := range cases {
		set := NewSet(c.input)
		if set != c.expected {
			t.Errorf("NewSet(%q) == %+v, expected %+v", c.input, set, c.expected)
		}
	}
}

func TestNewGame(t *testing.T) {
	cases := []struct {
		input    string
		expected Game
	}{
		{
			"Game 1: 4 red, 3 blue; 6 blue, 16 green; 9 blue, 13 green, 1 red; 10 green, 4 red, 6 blue\n",
			Game{1, []Set{{4, 0, 3}, {0, 16, 6}, {1, 13, 9}, {4, 10, 6}}},
		},
		{
			"Game 100: 7 blue, 6 red, 5 green; 3 blue, 13 green, 11 red; 6 red, 13 green, 14 blue; 8 red, 10 blue, 15 green",
			Game{100, []Set{{6, 5, 7}, {11, 13, 3}, {6, 13, 14}, {8, 15, 10}}},
		},
	}
	for _, c := range cases {
		game := NewGame(c.input)
		same := game.Id == c.expected.Id && len(game.Sets) == len(c.expected.Sets)
		if same {
			for i, set := range game.Sets {
				if set != c.expected.Sets[i] {
					same = false
					break
				}
			}
		}
		if !same {
			t.Errorf("NewGame(%q) == %+v, expected %+v", c.input, game, c.expected)
		}
	}
}

func TestSetCheck(t *testing.T) {
	cases := []struct {
		input    Set
		expected bool
	}{
		{Set{4, 0, 3}, true},
		{Set{4, 10, 6}, true},
		{Set{11, 0, 0}, true},
		{Set{20, 8, 6}, false},
		{Set{14, 3, 15}, false},
		{Set{4, 14, 3}, false},
	}
	for _, c := range cases {
		check := c.input.Check()
		if check != c.expected {
			t.Errorf("set.Check(%+v) == %t, expected %t", c.input, check, c.expected)
		}
	}
}

func TestSetPower(t *testing.T) {
	cases := []struct {
		input    Set
		expected int
	}{
		{Set{4, 0, 3}, 0},
		{Set{4, 10, 6}, 240},
		{Set{11, 0, 0}, 0},
		{Set{20, 8, 6}, 960},
		{Set{14, 3, 15}, 630},
		{Set{4, 14, 3}, 168},
	}
	for _, c := range cases {
		power := c.input.Power()
		if power != c.expected {
			t.Errorf("set.Power(%+v) == %d, expected %d", c.input, power, c.expected)
		}
	}
}

func TestSetMergeMax(t *testing.T) {
	cases := []struct {
		input    Set
		other    Set
		expected Set
	}{
		{Set{4, 0, 3}, Set{4, 10, 6}, Set{4, 10, 6}},
		{Set{11, 0, 0}, Set{20, 8, 6}, Set{20, 8, 6}},
		{Set{14, 3, 15}, Set{4, 14, 3}, Set{14, 14, 15}},
	}
	for _, c := range cases {
		set := c.input.MergeMax(c.other)
		if set != c.expected {
			t.Errorf("%#v.MergeMax(%+v) == %+v, expected %+v", c.input, c.other, set, c.expected)
		}
	}
}

func TestGameCheck(t *testing.T) {
	cases := []struct {
		input    Game
		expected bool
	}{
		{Game{1, []Set{{4, 0, 3}, {0, 16, 6}, {1, 13, 9}, {4, 10, 6}}}, false},
		{Game{100, []Set{{6, 5, 7}, {11, 13, 3}, {6, 13, 14}, {8, 15, 10}}}, false},
		{Game{1, []Set{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}}, true},
		{Game{2, []Set{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}}, true},
		{Game{3, []Set{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}}}, false},
		{Game{4, []Set{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}}}, false},
		{Game{5, []Set{{6, 3, 1}, {1, 2, 2}}}, true},
	}
	for _, c := range cases {
		check := c.input.Check()
		if check != c.expected {
			t.Errorf("game.Check(%+v) == %t, expected %t", c.input, check, c.expected)
		}
	}
}

func TestGameLeastUpperBound(t *testing.T) {
	cases := []struct {
		input    Game
		expected Set
	}{
		{Game{1, []Set{{4, 0, 3}, {0, 16, 6}, {1, 13, 9}, {4, 10, 6}}}, Set{4, 16, 9}},
		{Game{100, []Set{{6, 5, 7}, {11, 13, 3}, {6, 13, 14}, {8, 15, 10}}}, Set{11, 15, 14}},
		{Game{1, []Set{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}}, Set{4, 2, 6}},
		{Game{2, []Set{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}}, Set{1, 3, 4}},
		{Game{3, []Set{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}}}, Set{20, 13, 6}},
		{Game{4, []Set{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}}}, Set{14, 3, 15}},
		{Game{5, []Set{{6, 3, 1}, {1, 2, 2}}}, Set{6, 3, 2}},
	}
	for _, c := range cases {
		lub := c.input.LeastUpperBound()
		if lub != c.expected {
			t.Errorf("game.LeastUpperBound(%+v) == %+v, expected %+v", c.input, lub, c.expected)
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
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			8,
		},
	}
	for _, c := range cases {
		sum := Sum(c.inputs)
		if sum != c.expected {
			t.Errorf("Sum(%v) == %d, expected %d", c.inputs, sum, c.expected)
		}
	}
}

func TestSumPower(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			2286,
		},
	}
	for _, c := range cases {
		sum := SumPower(c.inputs)
		if sum != c.expected {
			t.Errorf("SumPower(%v) == %d, expected %d", c.inputs, sum, c.expected)
		}
	}
}
