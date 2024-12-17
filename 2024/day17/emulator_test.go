package day17

import "testing"

func TestExecProgram(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected string
	}{
		{
			inputs: []string{
				"Register A: 10",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 5,0,5,1,5,4",
			},
			expected: "0,1,2",
		},
		{
			inputs: []string{
				"Register A: 2024",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			expected: "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			inputs: []string{
				"Register A: 729",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			expected: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			inputs: []string{
				"Register A: 0",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "4",
		},
		{
			inputs: []string{
				"Register A: 1",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "4",
		},
		{
			inputs: []string{
				"Register A: 2",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "6",
		},
		{
			inputs: []string{
				"Register A: 3",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "7",
		},
		{
			inputs: []string{
				"Register A: 4",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "0",
		},
		{
			inputs: []string{
				"Register A: 5",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "1",
		},
		{
			inputs: []string{
				"Register A: 6",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "2",
		},
		{
			inputs: []string{
				"Register A: 7",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "3",
		},
		{
			inputs: []string{
				"Register A: 8",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			},
			expected: "0,4",
		},
	}
	for _, c := range cases {
		result := ExecProgram(c.inputs)
		if result != c.expected {
			t.Errorf("ExecProgram(%q) == %q, expected %q", c.inputs, result, c.expected)
		}
	}
}
