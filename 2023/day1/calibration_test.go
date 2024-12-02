package day1

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foobar", "raboof"},
		{"crash", "hsarc"},
	}
	for _, c := range cases {
		result := Reverse(c.input)
		if result != c.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", c.input, result, c.expected)
		}
	}
}

func TestFindValue(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"zlmlk1", 11},
		{"vqjvxtc79mvdnktdsxcqc1sevenone", 71},
		{"9kkpjrmhmmlzvqngnhftwoonexjqmnfive", 95},
		{"sevenone", 71},
		{"7pqrst6teen", 76},
		{"fivetczxxvjrrqfive1sevennvj6one3", 53},
		{"22pztfnhh554qdvzjxnnzzrj", 24},
		{"one", 11},
		{"two", 22},
		{"three", 33},
		{"four", 44},
		{"five", 55},
		{"six", 66},
		{"seven", 77},
		{"eight", 88},
		{"nine", 99},
		{"1", 11},
		{"2", 22},
		{"3", 33},
		{"4", 44},
		{"5", 55},
		{"6", 66},
		{"7", 77},
		{"8", 88},
		{"9", 99},
		{"twone", 21},
		{"4twonel", 41},
		{"jgtwonetwosixthreervlmxlnine869lbqzxpqqn", 29},
		{"99dndgmkcctwoneftn", 91},
		{"eightwo", 82},
		{"248twofbkfpxtheightwovng", 22},
		{"deightwoeighteight5", 85},
		{"ninesixpfrjvfm8kkjsrhttxsslhtwoeightwovvg", 92},
		{"oneight", 18},
		{"threeight", 38},
		{"fiveight", 58},
		{"nineight", 98},
		{"sevenine", 79},
		{"eighthree", 83},
	}
	for _, c := range cases {
		result := FindValue(c.input)
		if result != c.expected {
			t.Errorf("FindValue(%q) == %d, expected %d", c.input, result, c.expected)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{
			[]string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			142,
		},
		{
			[]string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			281,
		},
	}

	for _, c := range cases {
		result := Sum(c.inputs)
		if result != c.expected {
			t.Errorf("Sum(%v) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}
