package day5part2

import (
	"testing"
)

func TestApply(t *testing.T) {
	cases := []struct {
		rm   RangeMap
		rs   []Range
		want []Range
	}{
		{
			rm:   RangeMap{{50, 98, 2}, {52, 50, 48}},
			rs:   []Range{{79, 14}, {55, 13}},
			want: []Range{{81, 14}, {57, 13}},
		},
		{
			rm:   RangeMap{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}},
			rs:   []Range{{81, 14}, {57, 13}},
			want: []Range{{81, 14}, {57, 13}},
		},
		{
			rm:   RangeMap{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}},
			rs:   []Range{{81, 14}, {57, 13}},
			want: []Range{{81, 14}, {53, 4}, {61, 9}},
		},
	}
	for _, c := range cases {
		got := c.rm.Apply(c.rs)
		if !got.Equal(c.want) {
			t.Errorf("RangeMap.Apply(%v) == %v, want %v", c.rs, got, c.want)
		}
	}
}

func TestMinLocation(t *testing.T) {
	cases := []struct {
		lines []string
		want  int
	}{
		{
			lines: []string{
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
			want: 46,
		},
	}
	for _, c := range cases {
		got := MinLocation(c.lines)
		if got != c.want {
			t.Errorf("MinLocation(%v) == %d, want %d", c.lines, got, c.want)
		}
	}
}
