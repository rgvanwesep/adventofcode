package day5

import (
	"reflect"
	"testing"
)

func TestNewAtlas(t *testing.T) {
	cases := []struct {
		inputs []string
		atlas  Atlas
	}{
		{
			inputs: []string{
				"seeds: 79 14 55 13\n",
				"\n",
				"seed-to-soil map:\n",
				"50 98 2\n",
				"52 50 48\n",
				"\n",
				"soil-to-fertilizer map:\n",
				"0 15 37\n",
				"37 52 2\n",
				"39 0 15\n",
				"\n",
				"fertilizer-to-water map:\n",
				"49 53 8\n",
				"0 11 42\n",
				"42 0 7\n",
				"57 7 4\n",
				"\n",
				"water-to-light map:\n",
				"88 18 7\n",
				"18 25 70\n",
				"\n",
				"light-to-temperature map:\n",
				"45 77 23\n",
				"81 45 19\n",
				"68 64 13\n",
				"\n",
				"temperature-to-humidity map:\n",
				"0 69 1\n",
				"1 0 69\n",
				"\n",
				"humidity-to-location map:\n",
				"60 56 37\n",
				"56 93 4",
			},
			atlas: Atlas{
				Seeds: Seeds{79, 14, 55, 13},
				Maps: []RangeMap{
					{
						Header: "seed-to-soil map:",
						Items: []RangeMapItem{
							{50, 98, 2},
							{52, 50, 48},
						},
					},
					{
						Header: "soil-to-fertilizer map:",
						Items: []RangeMapItem{
							{0, 15, 37},
							{37, 52, 2},
							{39, 0, 15},
						},
					},
					{
						Header: "fertilizer-to-water map:",
						Items: []RangeMapItem{
							{49, 53, 8},
							{0, 11, 42},
							{42, 0, 7},
							{57, 7, 4},
						},
					},
					{
						Header: "water-to-light map:",
						Items: []RangeMapItem{
							{88, 18, 7},
							{18, 25, 70},
						},
					},
					{
						Header: "light-to-temperature map:",
						Items: []RangeMapItem{
							{45, 77, 23},
							{81, 45, 19},
							{68, 64, 13},
						},
					},
					{
						Header: "temperature-to-humidity map:",
						Items: []RangeMapItem{
							{0, 69, 1},
							{1, 0, 69},
						},
					},
					{
						Header: "humidity-to-location map:",
						Items: []RangeMapItem{
							{60, 56, 37},
							{56, 93, 4},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		atlas := NewAtlas(c.inputs)
		if !reflect.DeepEqual(atlas, c.atlas) {
			t.Errorf("NewAtlas(%v) == %v, want %v", c.inputs, atlas, c.atlas)
		}
	}
}

func TestItemApply(t *testing.T) {
	cases := []struct {
		item     RangeMapItem
		input    int
		expected int
		ok       bool
	}{
		{
			RangeMapItem{50, 98, 2},
			98,
			50,
			true,
		},
		{
			RangeMapItem{50, 98, 2},
			99,
			51,
			true,
		},
		{
			RangeMapItem{50, 98, 2},
			97,
			97,
			false,
		},
		{
			RangeMapItem{52, 50, 48},
			100,
			100,
			false,
		},
	}
	for _, c := range cases {
		actual, ok := c.item.Apply(c.input)
		if ok != c.ok || actual != c.expected {
			t.Errorf("Apply(%v, %v) == %v, want %v", c.item, c.input, actual, c.expected)
		}
	}
}

func TestRangeMapApply(t *testing.T) {
	cases := []struct {
		r        RangeMap
		input    int
		expected []int
	}{
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			98,
			[]int{50},
		},
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			99,
			[]int{51},
		},
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			97,
			[]int{99},
		},
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			100,
			[]int{100},
		},
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			50,
			[]int{52},
		},
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			53,
			[]int{55},
		},
		{
			RangeMap{
				Header: "seed-to-soil map:",
				Items: []RangeMapItem{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			49,
			[]int{49},
		},
	}
	for _, c := range cases {
		actual := c.r.Apply(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("Apply(%v, %v) == %v, want %v", c.r, c.input, actual, c.expected)
		}
	}
}

func TestFindLocations(t *testing.T) {
	cases := []struct {
		atlas    Atlas
		expected []int
	}{
		{
			Atlas{
				Seeds: Seeds{79, 14, 55, 13},
				Maps: []RangeMap{
					{
						Header: "seed-to-soil map:",
						Items: []RangeMapItem{
							{50, 98, 2},
							{52, 50, 48},
						},
					},
					{
						Header: "soil-to-fertilizer map:",
						Items: []RangeMapItem{
							{0, 15, 37},
							{37, 52, 2},
							{39, 0, 15},
						},
					},
					{
						Header: "fertilizer-to-water map:",
						Items: []RangeMapItem{
							{49, 53, 8},
							{0, 11, 42},
							{42, 0, 7},
							{57, 7, 4},
						},
					},
					{
						Header: "water-to-light map:",
						Items: []RangeMapItem{
							{88, 18, 7},
							{18, 25, 70},
						},
					},
					{
						Header: "light-to-temperature map:",
						Items: []RangeMapItem{
							{45, 77, 23},
							{81, 45, 19},
							{68, 64, 13},
						},
					},
					{
						Header: "temperature-to-humidity map:",
						Items: []RangeMapItem{
							{0, 69, 1},
							{1, 0, 69},
						},
					},
					{
						Header: "humidity-to-location map:",
						Items: []RangeMapItem{
							{60, 56, 37},
							{56, 93, 4},
						},
					},
				},
			},
			[]int{82, 43, 86, 35},
		},
	}
	for _, c := range cases {
		actual := c.atlas.FindLocations()
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("FindLocations(%v) == %v, want %v", c.atlas, actual, c.expected)
		}
	}
}

func TestMinLocation(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"seeds: 79 14 55 13\n",
				"\n",
				"seed-to-soil map:\n",
				"50 98 2\n",
				"52 50 48\n",
				"\n",
				"soil-to-fertilizer map:\n",
				"0 15 37\n",
				"37 52 2\n",
				"39 0 15\n",
				"\n",
				"fertilizer-to-water map:\n",
				"49 53 8\n",
				"0 11 42\n",
				"42 0 7\n",
				"57 7 4\n",
				"\n",
				"water-to-light map:\n",
				"88 18 7\n",
				"18 25 70\n",
				"\n",
				"light-to-temperature map:\n",
				"45 77 23\n",
				"81 45 19\n",
				"68 64 13\n",
				"\n",
				"temperature-to-humidity map:\n",
				"0 69 1\n",
				"1 0 69\n",
				"\n",
				"humidity-to-location map:\n",
				"60 56 37\n",
				"56 93 4",
			},
			expected: 35,
		},
	}
	for _, c := range cases {
		actual := MinLocation(c.inputs)
		if actual != c.expected {
			t.Errorf("MinLocation(%v) == %v, want %v", c.inputs, actual, c.expected)
		}
	}
}
