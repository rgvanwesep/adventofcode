package day4

import "testing"

func TestAdd(t *testing.T) {
	set := make(Set)
	set.Add(1)
	if !set[1] {
		t.Error("Set should contain 1")
	}
}

func TestIntersect(t *testing.T) {
	cases := []struct {
		a        Set
		b        Set
		expected Set
	}{
		{
			a:        Set{1: true, 2: true, 3: true},
			b:        Set{1: true, 2: true, 3: true},
			expected: Set{1: true, 2: true, 3: true},
		},
		{
			a:        Set{1: true, 2: true, 3: true},
			b:        Set{1: true, 2: true},
			expected: Set{1: true, 2: true},
		},
		{
			a:        Set{1: true, 2: true},
			b:        Set{1: true, 2: true, 3: true},
			expected: Set{1: true, 2: true},
		},
		{
			a:        Set{1: true, 2: true, 3: true},
			b:        Set{4: true, 5: true, 6: true},
			expected: Set{},
		},
		{
			a:        Set{1: true, 2: true, 3: true},
			b:        Set{},
			expected: Set{},
		},
		{
			a:        Set{},
			b:        Set{1: true, 2: true, 3: true},
			expected: Set{},
		},
	}
	for _, c := range cases {
		actual := c.a.Intersect(c.b)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, actual)
		}
		for k := range c.expected {
			if !actual[k] {
				t.Errorf("Expected %v, got %v", c.expected, actual)
			}
		}
	}
}

func TestNewCard(t *testing.T) {
	cases := []struct {
		input    string
		expected Card
	}{
		{
			"Card   1: 34 50 18 44 19 35 47 62 65 26 | 63  6 27 15 60  9 98  3 61 89 31 43 80 37 54 49 92 55  8  7 10 16 52 33 45",
			Card{
				Id:      1,
				Winners: Set{34: true, 50: true, 18: true, 44: true, 19: true, 35: true, 47: true, 62: true, 65: true, 26: true},
				Picks: Set{
					63: true, 6: true, 27: true, 15: true, 60: true, 9: true, 98: true, 3: true, 61: true, 89: true,
					31: true, 43: true, 80: true, 37: true, 54: true, 49: true, 92: true, 55: true, 8: true, 7: true, 10: true,
					16: true, 52: true, 33: true, 45: true,
				},
			},
		},
		{
			"Card  34:  8 77 92 34 84 28 90 40 97 75 | 61 40 99 77 17 28 80 50 37 47 22 70 81 79 97 85 93 15 49 48 69 14  2 12 94",
			Card{
				Id:      34,
				Winners: Set{8: true, 77: true, 92: true, 34: true, 84: true, 28: true, 90: true, 40: true, 97: true, 75: true},
				Picks: Set{
					61: true, 40: true, 99: true, 77: true, 17: true, 28: true, 80: true, 50: true, 37: true, 47: true,
					22: true, 70: true, 81: true, 79: true, 97: true, 85: true, 93: true, 15: true, 49: true, 48: true, 69: true,
					14: true, 2: true, 12: true, 94: true,
				},
			},
		},
		{
			"Card  51: 22 94 42 24 28 37 61 88 86 12 |  5 31  3 34 56 82 70 68 39 91 53 22 16 81 71 54 99 41 44 90 24 37 12 27 61",
			Card{
				Id:      51,
				Winners: Set{22: true, 94: true, 42: true, 24: true, 28: true, 37: true, 61: true, 88: true, 86: true, 12: true},
				Picks: Set{
					5: true, 31: true, 3: true, 34: true, 56: true, 82: true, 70: true, 68: true, 39: true, 91: true,
					53: true, 22: true, 16: true, 81: true, 71: true, 54: true, 99: true, 41: true, 44: true, 90: true, 24: true,
					37: true, 12: true, 27: true, 61: true,
				},
			},
		},
		{
			"Card 142: 34 71 94  2 79 18 69 89 44 19 |  3 10  9 62 71 44 37 32 97 85  2 89 48  6 14 95 17 91  5 99 11 33 41 39 22",
			Card{
				Id:      142,
				Winners: Set{34: true, 71: true, 94: true, 2: true, 79: true, 18: true, 69: true, 89: true, 44: true, 19: true},
				Picks: Set{
					3: true, 10: true, 9: true, 62: true, 71: true, 44: true, 37: true, 32: true, 97: true, 85: true,
					2: true, 89: true, 48: true, 6: true, 14: true, 95: true, 17: true, 91: true, 5: true, 99: true, 11: true,
					33: true, 41: true, 39: true, 22: true,
				},
			},
		},
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			Card{
				Id:      1,
				Winners: Set{41: true, 48: true, 83: true, 86: true, 17: true},
				Picks:   Set{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
			},
		},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			Card{
				Id:      2,
				Winners: Set{13: true, 32: true, 20: true, 16: true, 61: true},
				Picks:   Set{61: true, 30: true, 68: true, 82: true, 17: true, 32: true, 24: true, 19: true},
			},
		},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			Card{
				Id:      3,
				Winners: Set{1: true, 21: true, 53: true, 59: true, 44: true},
				Picks:   Set{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 1: true},
			},
		},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			Card{
				Id:      4,
				Winners: Set{41: true, 92: true, 73: true, 84: true, 69: true},
				Picks:   Set{59: true, 84: true, 76: true, 51: true, 58: true, 5: true, 54: true, 83: true},
			},
		},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			Card{
				Id:      5,
				Winners: Set{87: true, 83: true, 26: true, 28: true, 32: true},
				Picks:   Set{88: true, 30: true, 70: true, 12: true, 93: true, 22: true, 82: true, 36: true},
			},
		},
		{
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			Card{
				Id:      6,
				Winners: Set{31: true, 18: true, 13: true, 56: true, 72: true},
				Picks:   Set{74: true, 77: true, 10: true, 23: true, 35: true, 67: true, 36: true, 11: true},
			},
		},
	}
	for _, c := range cases {
		actual := NewCard(c.input)
		if actual.Id != c.expected.Id {
			t.Errorf("Expected %v, got %v", c.expected.Id, actual.Id)
		}
		if len(actual.Winners) != len(c.expected.Winners) {
			t.Errorf("Expected %v, got %v", c.expected.Winners, actual.Winners)
		}
		for k := range c.expected.Winners {
			if !actual.Winners[k] {
				t.Errorf("Expected %v, got %v", c.expected.Winners, actual.Winners)
			}
		}
		if len(actual.Picks) != len(c.expected.Picks) {
			t.Errorf("Expected %v, got %v", c.expected.Picks, actual.Picks)
		}
		for k := range c.expected.Picks {
			if !actual.Picks[k] {
				t.Errorf("Expected %v, got %v", c.expected.Picks, actual.Picks)
			}
		}
	}
}

func TestScore(t *testing.T) {
	cases := []struct {
		card     Card
		expected int
	}{
		{
			Card{
				Id:      1,
				Winners: Set{},
				Picks:   Set{},
			},
			0,
		},
		{
			Card{
				Id:      1,
				Winners: Set{1: true, 2: true, 3: true},
				Picks:   Set{},
			},
			0,
		},
		{
			Card{
				Id:      1,
				Winners: Set{},
				Picks:   Set{1: true, 2: true, 3: true},
			},
			0,
		},
		{
			Card{
				Id:      1,
				Winners: Set{1: true, 2: true, 3: true},
				Picks:   Set{1: true, 2: true, 3: true},
			},
			4,
		},
		{
			Card{
				Id:      1,
				Winners: Set{41: true, 48: true, 83: true, 86: true, 17: true},
				Picks:   Set{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
			},
			8,
		},
		{
			Card{
				Id:      2,
				Winners: Set{13: true, 32: true, 20: true, 16: true, 61: true},
				Picks:   Set{61: true, 30: true, 68: true, 82: true, 17: true, 32: true, 24: true, 19: true},
			},
			2,
		},
		{
			Card{
				Id:      3,
				Winners: Set{1: true, 21: true, 53: true, 59: true, 44: true},
				Picks:   Set{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 1: true},
			},
			2,
		},
		{
			Card{
				Id:      4,
				Winners: Set{41: true, 92: true, 73: true, 84: true, 69: true},
				Picks:   Set{59: true, 84: true, 76: true, 51: true, 58: true, 5: true, 54: true, 83: true},
			},
			1,
		},
		{
			Card{
				Id:      5,
				Winners: Set{87: true, 83: true, 26: true, 28: true, 32: true},
				Picks:   Set{88: true, 30: true, 70: true, 12: true, 93: true, 22: true, 82: true, 36: true},
			},
			0,
		},
		{
			Card{
				Id:      6,
				Winners: Set{31: true, 18: true, 13: true, 56: true, 72: true},
				Picks:   Set{74: true, 77: true, 10: true, 23: true, 35: true, 67: true, 36: true, 11: true},
			},
			0,
		},
	}
	for _, c := range cases {
		actual := c.card.Score()
		if actual != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, actual)
		}
	}
}

func TestNumMatches(t *testing.T) {
	cases := []struct {
		card     Card
		expected int
	}{
		{
			Card{
				Id:      1,
				Winners: Set{},
				Picks:   Set{},
			},
			0,
		},
		{
			Card{
				Id:      1,
				Winners: Set{1: true, 2: true, 3: true},
				Picks:   Set{},
			},
			0,
		},
		{
			Card{
				Id:      1,
				Winners: Set{},
				Picks:   Set{1: true, 2: true, 3: true},
			},
			0,
		},
		{
			Card{
				Id:      1,
				Winners: Set{1: true, 2: true, 3: true},
				Picks:   Set{1: true, 2: true, 3: true},
			},
			3,
		},
	}
	for _, c := range cases {
		actual := c.card.NumMatches()
		if actual != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, actual)
		}
		actual = c.card.NumMatches()
		if actual != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, actual)
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
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			13,
		},
	}
	for _, c := range cases {
		actual := Sum(c.inputs)
		if actual != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, actual)
		}
	}
}

func TestSumCards(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			30,
		},
	}
	for _, c := range cases {
		actual := SumCards(c.inputs)
		if actual != c.expected {
			t.Errorf("Expected %v, got %v", c.expected, actual)
		}
	}
}
