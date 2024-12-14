package day13

import "testing"

func TestMinCost(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=94, Y=34",
			},
			3,
		},
		{
			[]string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=22, Y=67",
			},
			1,
		},
		{
			[]string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=116, Y=101",
			},
			4,
		},
		{
			[]string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=1160, Y=1010",
			},
			40,
		},
		{
			[]string{
				/*
					94a + 22b = 8400
					34a + 67b = 5400

					1598 = lcm(94, 34)
					17 = 1598/94
					47 = 1598/34

					1598a + 374b  = 142800
					1598a + 3149b = 253800

					(3149 - 374)b = 253800 - 142800
					2775b = 111000
					b = 40
					a = (142800 - 374*40)/1598
					a = 80

					c = 3*80 + 40
					c = 280

				*/
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=8400, Y=5400",
				"",
				"Button A: X+26, Y+66",
				"Button B: X+67, Y+21",
				"Prize: X=12748, Y=12176",
				"",
				"Button A: X+17, Y+86",
				"Button B: X+84, Y+37",
				"Prize: X=7870, Y=6450",
				"",
				"Button A: X+69, Y+23",
				"Button B: X+27, Y+71",
				"Prize: X=18641, Y=10279",
			},
			480,
		},
	}
	for _, c := range cases {
		result := MinCost(c.inputs)
		if result != c.expected {
			t.Errorf("MinCost(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
