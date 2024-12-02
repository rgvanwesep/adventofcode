package day8

import "testing"

func TestNewNode(t *testing.T) {
	cases := []struct {
		line string
		want Node
	}{
		{
			"FJT = (XDJ, LQV)\n",
			Node{
				Id:    "FJT",
				Left:  "XDJ",
				Right: "LQV",
			},
		},
		{
			"VMG = (DNX, BDL)",
			Node{
				Id:    "VMG",
				Left:  "DNX",
				Right: "BDL",
			},
		},
	}
	for _, c := range cases {
		got, _ := NewNode(c.line)
		if got != c.want {
			t.Errorf("NewNode(%q) == %v, want %v", c.line, got, c.want)
		}
	}
}

func TestNewDirections(t *testing.T) {
	cases := []struct {
		line string
		want Directions
	}{
		{
			"LLR",
			Directions("LLR"),
		},
		{
			"RRR",
			Directions("RRR"),
		},
	}
	for _, c := range cases {
		got, _ := NewDirections(c.line)
		if got != c.want {
			t.Errorf("NewDirections(%q) == %v, want %v", c.line, got, c.want)
		}
	}
}

func TestNewGraph(t *testing.T) {
	cases := []struct {
		lines []string
		want  Graph
	}{
		{
			[]string{
				"FJT = (XDJ, LQV)\n",
				"VMG = (DNX, BDL)",
			},
			Graph{
				Nodes: map[string]Node{
					"FJT": {
						Id:    "FJT",
						Left:  "XDJ",
						Right: "LQV",
					},
					"VMG": {
						Id:    "VMG",
						Left:  "DNX",
						Right: "BDL",
					},
				},
			},
		},
		{
			[]string{
				"\n",
				"AAA = (BBB, CCC)\n",
				"BBB = (DDD, EEE)\n",
				"CCC = (ZZZ, GGG)\n",
				"DDD = (DDD, DDD)\n",
				"EEE = (EEE, EEE)\n",
				"GGG = (GGG, GGG)\n",
				"ZZZ = (ZZZ, ZZZ)",
			},
			Graph{
				Nodes: map[string]Node{
					"AAA": {
						Id:    "AAA",
						Left:  "BBB",
						Right: "CCC",
					},
					"BBB": {
						Id:    "BBB",
						Left:  "DDD",
						Right: "EEE",
					},
					"CCC": {
						Id:    "CCC",
						Left:  "ZZZ",
						Right: "GGG",
					},
					"DDD": {
						Id:    "DDD",
						Left:  "DDD",
						Right: "DDD",
					},
					"EEE": {
						Id:    "EEE",
						Left:  "EEE",
						Right: "EEE",
					},
					"GGG": {
						Id:    "GGG",
						Left:  "GGG",
						Right: "GGG",
					},
					"ZZZ": {
						Id:    "ZZZ",
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
		{
			[]string{
				"AAA = (BBB, BBB)\n",
				"BBB = (AAA, ZZZ)\n",
				"ZZZ = (ZZZ, ZZZ)",
			},
			Graph{
				Nodes: map[string]Node{
					"AAA": {
						Id:    "AAA",
						Left:  "BBB",
						Right: "BBB",
					},
					"BBB": {
						Id:    "BBB",
						Left:  "AAA",
						Right: "ZZZ",
					},
					"ZZZ": {
						Id:    "ZZZ",
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := NewGraph(c.lines)
		if !got.Equals(c.want) {
			t.Errorf("NewGraph(%q) == %v, want %v", c.lines, got, c.want)
		}
	}
}

func TestCountSteps(t *testing.T) {
	cases := []struct {
		lines []string
		want  int
	}{
		{
			[]string{
				"RL\n",
				"\n",
				"AAA = (BBB, CCC)\n",
				"BBB = (DDD, EEE)\n",
				"CCC = (ZZZ, GGG)\n",
				"DDD = (DDD, DDD)\n",
				"EEE = (EEE, EEE)\n",
				"GGG = (GGG, GGG)\n",
				"ZZZ = (ZZZ, ZZZ)",
			},
			2,
		},
		{
			[]string{
				"LLR\n",
				"\n",
				"AAA = (BBB, BBB)\n",
				"BBB = (AAA, ZZZ)\n",
				"ZZZ = (ZZZ, ZZZ)",
			},
			6,
		},
	}
	for _, c := range cases {
		got := CountSteps(c.lines)
		if got != c.want {
			t.Errorf("CountSteps(%q) == %v, want %v", c.lines, got, c.want)
		}
	}
}

func TestCountParallelSteps(t *testing.T) {
	cases := []struct {
		lines []string
		want  int
	}{
		{
			[]string{
				"LR\n",
				"\n",
				"11A = (11B, XXX)\n",
				"11B = (XXX, 11Z)\n",
				"11Z = (11B, XXX)\n",
				"22A = (22B, XXX)\n",
				"22B = (22C, 22C)\n",
				"22C = (22Z, 22Z)\n",
				"22Z = (22B, 22B)\n",
				"XXX = (XXX, XXX)",
			},
			6,
		},
	}
	for _, c := range cases {
		got := CountParallelSteps(c.lines)
		if got != c.want {
			t.Errorf("CountSteps(%q) == %v, want %v", c.lines, got, c.want)
		}
	}
}
