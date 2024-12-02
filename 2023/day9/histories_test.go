package day9

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		lines []string
		want  int
	}{
		{
			[]string{
				"0 3 6 9 12 15\n",
				"1 3 6 10 15 21\n",
				"10 13 16 21 30 45",
			},
			114,
		},
	}
	for _, c := range cases {
		got := Sum(c.lines)
		if got != c.want {
			t.Errorf("Sum(%v) == %d, want %d", c.lines, got, c.want)
		}
	}
}

func TestSumPrev(t *testing.T) {
	cases := []struct {
		lines []string
		want  int
	}{
		{
			[]string{
				"0 3 6 9 12 15\n",
				"1 3 6 10 15 21\n",
				"10 13 16 21 30 45",
			},
			2,
		},
	}
	for _, c := range cases {
		got := SumPrev(c.lines)
		if got != c.want {
			t.Errorf("SumPrev(%v) == %d, want %d", c.lines, got, c.want)
		}
	}
}
