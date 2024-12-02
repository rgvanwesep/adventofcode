package day7

import "testing"

func TestWinnings(t *testing.T) {
	cases := []struct {
		lines    []string
		expected int
	}{
		{
			[]string{
				"32T3K 765\n",
				"T55J5 684\n",
				"KK677 28\n",
				"KTJJT 220\n",
				"QQQJA 483",
			},
			6440,
		},
	}
	for _, c := range cases {
		winnings := Winnings(c.lines)
		if winnings != c.expected {
			t.Errorf("Winnings(%v) == %v, expected %v", c.lines, winnings, c.expected)
		}
	}
}

func TestJokerType(t *testing.T) {
	cases := []struct {
		hand     Hand
		expected Type
	}{
		{
			Hand{Cards: [handSize]Card{'3', '2', 'T', '3', 'K'}},
			onePair,
		},
		{
			Hand{Cards: [handSize]Card{'T', '5', '5', 'J', '5'}},
			fourOfAKind,
		},
		{
			Hand{Cards: [handSize]Card{'K', 'K', '6', '7', '7'}},
			twoPair,
		},
		{
			Hand{Cards: [handSize]Card{'K', 'T', 'J', 'J', 'T'}},
			fourOfAKind,
		},
		{
			Hand{Cards: [handSize]Card{'Q', 'Q', 'Q', 'J', 'A'}},
			fourOfAKind,
		},
	}
	for _, c := range cases {
		jokerType := c.hand.JokerType()
		if jokerType != c.expected {
			t.Errorf("JokerType(%v) == %v, expected %v", c.hand, jokerType, c.expected)
		}
	}
}

func TestJokerWinnings(t *testing.T) {
	cases := []struct {
		lines    []string
		expected int
	}{
		{
			[]string{
				"32T3K 765\n",
				"T55J5 684\n",
				"KK677 28\n",
				"KTJJT 220\n",
				"QQQJA 483",
			},
			5905,
		},
	}
	for _, c := range cases {
		winnings := JokerWinnings(c.lines)
		if winnings != c.expected {
			t.Errorf("Winnings(%v) == %v, expected %v", c.lines, winnings, c.expected)
		}
	}
}
