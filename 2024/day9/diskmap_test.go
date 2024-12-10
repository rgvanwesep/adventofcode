package day9

import "testing"

func TestCalcChecksum(t *testing.T) {
	cases := []struct {
		inputs   []string
		expected int
	}{
		{[]string{"12345"}, 60},
		{[]string{"2333133121414131402"}, 1928},
	}
	for _, c := range cases {
		result := CalcChecksum(c.inputs)
		if result != c.expected {
			t.Errorf("CalcChecksum(%q) == %d, expected %d", c.inputs, result, c.expected)
		}
	}
}

func TestExpandDiskMap(t *testing.T) {
	cases := []struct {
		input    string
		expected ExpandedDiskMap
	}{
		{
			"12345",
			ExpandedDiskMap{
				blocks: []Block{
					{empty: false, fileId: 0},
					{empty: true, fileId: -1},
					{empty: true, fileId: -1},
					{empty: false, fileId: 1},
					{empty: false, fileId: 1},
					{empty: false, fileId: 1},
					{empty: true, fileId: -1},
					{empty: true, fileId: -1},
					{empty: true, fileId: -1},
					{empty: true, fileId: -1},
					{empty: false, fileId: 2},
					{empty: false, fileId: 2},
					{empty: false, fileId: 2},
					{empty: false, fileId: 2},
					{empty: false, fileId: 2},
				},
				emptyIndices:  []int{1, 2, 6, 7, 8, 9},
				lastFileIndex: 14,
			},
		},
	}
	for _, c := range cases {
		result := ExpandDiskMap(c.input)
		mismatched := false
		if len(result.blocks) != len(c.expected.blocks) {
			mismatched = true
		}
		for i, b := range result.blocks {
			if b != c.expected.blocks[i] {
				mismatched = true
				break
			}
		}
		if len(result.emptyIndices) != len(c.expected.emptyIndices) {
			mismatched = true
		}
		for i, e := range result.emptyIndices {
			if e != c.expected.emptyIndices[i] {
				mismatched = true
			}
		}
		if result.lastFileIndex != c.expected.lastFileIndex {
			mismatched = true
		}
		if mismatched {
			t.Errorf("ExpandDiskMap(%q) == %v, expected %v", c.input, result, c.expected)
		}
	}
}
