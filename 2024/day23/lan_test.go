package day23

import "testing"

func TestCountLANs(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"kh-tc",
				"qp-kh",
				"de-cg",
				"ka-co",
				"yn-aq",
				"qp-ub",
				"cg-tb",
				"vc-aq",
				"tb-ka",
				"wh-tc",
				"yn-cg",
				"kh-ub",
				"ta-co",
				"de-co",
				"tc-td",
				"tb-wq",
				"wh-td",
				"ta-ka",
				"td-qp",
				"aq-cg",
				"wq-ub",
				"ub-vc",
				"de-ta",
				"wq-aq",
				"wq-vc",
				"wh-yn",
				"ka-de",
				"kh-ta",
				"co-tc",
				"wh-qp",
				"tb-vc",
				"td-yn",
			},
			expected: 7,
		},
	}
	for _, c := range cases {
		result := CountLANs(c.inputs)
		if result != c.expected {
			t.Errorf("CountLANs(%q) == %d, expected %d",
				c.inputs, result, c.expected,
			)
		}
	}
}
