package day6

import (
	"reflect"
	"testing"
)

func TestCanWin(t *testing.T) {
	race := Race{Time: 7, Distance: 9}
	cases := []struct {
		impulse  uint64
		expected bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, false},
		{7, false},
	}
	for _, c := range cases {
		canWin := race.CanWin(c.impulse)
		if canWin != c.expected {
			t.Errorf("Expected %v for impulse %v, got %v", c.expected, c.impulse, canWin)
		}
	}
}

func TestNumWins(t *testing.T) {
	cases := []struct {
		race     Race
		expected uint64
	}{
		{Race{7, 9}, 4},
		{Race{15, 40}, 8},
		{Race{30, 200}, 9},
	}
	for _, c := range cases {
		numWins := c.race.NumWins()
		if numWins != c.expected {
			t.Errorf("%v.NumWins() == %d, expected %d", c.race, numWins, c.expected)
		}
	}
}

func TestNewRaces(t *testing.T) {
	cases := []struct {
		inputs []string
		races  Races
	}{
		{
			[]string{
				"Time:      7  15   30\n",
				"Distance:  9  40  200",
			},
			Races{
				{7, 9},
				{15, 40},
				{30, 200},
			},
		},
	}
	for _, c := range cases {
		races := NewRaces(c.inputs)
		if !reflect.DeepEqual(races, c.races) {
			t.Errorf("NewRaces(%v) == %v, expected %v", c.inputs, races, c.races)
		}
	}
}

func TestProduct(t *testing.T) {
	cases := []struct {
		inputs  []string
		product uint64
	}{
		{
			[]string{
				"Time:      7  15   30\n",
				"Distance:  9  40  200",
			},
			288,
		},
	}
	for _, c := range cases {
		product := Product(c.inputs)
		if product != c.product {
			t.Errorf("Product(%v) == %d, expected %d", c.inputs, product, c.product)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		inputs []string
		count  uint64
	}{
		{
			[]string{
				"Time:      7  15   30\n",
				"Distance:  9  40  200",
			},
			71503,
		},
	}
	for _, c := range cases {
		count := Count(c.inputs)
		if count != c.count {
			t.Errorf("Count(%v) == %d, expected %d", c.inputs, count, c.count)
		}
	}
}
