package day2

import "testing"

func TestCountSafeReports(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected uint64
	}{
		{
			[]string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			2,
		},
		{
			[]string{
				"55 56 59 62 61",          // unsafe
				"68 70 71 74 75 76 78 78", // unsafe
				"52 55 56 58 62",          // unsafe
				"73 76 79 82 84 85 87 94", // unsafe
				"1 4 5 6 3 4",             // unsafe
				"77 80 78 80 83 81",       // unsafe
				"69 72 73 71 71",          // unsafe
				"14 15 17 18 17 21",       // unsafe
				"39 42 43 45 47 46 53",    // unsafe
				"72 73 73 74 75",          // unsafe
				"72 73 74 75",             // safe
				"72 72 74 75",             // unsafe
				"52 55 56 58 61",          // safe
				"100 97 94 91 88 85 82",   // safe
				"97 97 94 91 88 85 82",    // unsafe
				"100 97 94 91 88 85 85",   // unsafe
				"100 97 94 91 88 85 86",   // unsafe
			},
			3,
		},
	}
	for _, c := range cases {
		result := CountSafeReports(c.inputs)
		if result != c.expected {
			t.Errorf("CountSafeReports(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestCountSafeReportsDamped(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected uint64
	}{
		{
			[]string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			4,
		},
	}
	for _, c := range cases {
		result := CountSafeReportsDamped(c.inputs)
		if result != c.expected {
			t.Errorf("CountSafeReportsDamped(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
