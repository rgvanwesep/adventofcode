package day10

import "testing"

func TestSumTrailScores(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"0123",
				"1234",
				"8765",
				"9876",
			},
			1,
		},
		{
			[]string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			36,
		},
	}
	for _, c := range cases {
		result := SumTrailScores(c.inputs)
		if result != c.expected {
			t.Errorf("SumTrailScores(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestSumTrailRatings(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"0123",
				"1234",
				"8765",
				"9876",
			},
			16,
		},
		{
			[]string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			81,
		},
	}
	for _, c := range cases {
		result := SumTrailRatings(c.inputs)
		if result != c.expected {
			t.Errorf("SumTrailRatings(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
