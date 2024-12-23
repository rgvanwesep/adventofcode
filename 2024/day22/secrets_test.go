package day22

import "testing"

/*
1: 8685429
10: 4700978
100: 15273692
2024: 8667524
*/
func TestCalcFinalSecret(t *testing.T) {
	cases := []struct {
		seed     int
		nSecrets int
		expected int
	}{
		{
			seed:     1,
			nSecrets: 2000,
			expected: 8685429,
		},
		{
			seed:     10,
			nSecrets: 2000,
			expected: 4700978,
		},
		{
			seed:     2024,
			nSecrets: 2000,
			expected: 8667524,
		},
	}
	for _, c := range cases {
		result := calcFinalSecret(c.seed, c.nSecrets)
		if result != c.expected {
			t.Errorf("calcFinalSecret(%d, %d) == %d, expected %d",
				c.seed, c.nSecrets, result, c.expected,
			)
		}
	}
}
